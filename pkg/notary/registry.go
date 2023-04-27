package notary

import (
	"context"
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	gcrremote "github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/kyverno/kyverno/pkg/registryclient"
	notationregistry "github.com/notaryproject/notation-go/registry"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
	"oras.land/oras-go/v2/registry"
	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
)

func parseReference(ctx context.Context, ref string, registryClient registryclient.Client) (notationregistry.Repository, registry.Reference, error) {
	parsedRef, err := registry.ParseReference(ref)
	if err != nil {
		return nil, registry.Reference{}, errors.Wrapf(err, "failed to parse registry reference %s", ref)
	}

	authClient, plainHTTP, err := getAuthClient(ctx, parsedRef, registryClient)
	if err != nil {
		return nil, registry.Reference{}, err
	}

	repo, err := remote.NewRepository(ref)
	if err != nil {
		return nil, registry.Reference{}, errors.Wrapf(err, "failed to initialize repository")
	}

	repo.PlainHTTP = plainHTTP
	repo.Client = authClient
	repository := notationregistry.NewRepository(repo)

	parsedRef, err = resolveDigest(repository, parsedRef)
	if err != nil {
		return nil, registry.Reference{}, errors.Wrapf(err, "failed to resolve digest")
	}

	return repository, parsedRef, nil
}

func parseReferenceCrane(ctx context.Context, ref string, registryClient registryclient.Client) (notationregistry.Repository, crane.Option, []gcrremote.Option, error) {
	nameRef, err := name.ParseReference(ref)
	if err != nil {
		return nil, nil, nil, err
	}

	authenticator, err := getAuthenticator(ctx, ref, registryClient)
	if err != nil {
		return nil, nil, nil, err
	}
	craneOpts := crane.WithAuth(*authenticator)

	remoteOpts, err := getRemoteOpts(*authenticator)
	if err != nil {
		return nil, nil, nil, err
	}

	repository := NewRepository(craneOpts, remoteOpts, nameRef)

	err = resolveDigestCrane(repository, craneOpts, remoteOpts, nameRef)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "failed to resolve digest")
	}

	return repository, craneOpts, remoteOpts, nil
}

type imageResource struct {
	ref registry.Reference
}

func (ir *imageResource) String() string {
	return ir.ref.String()
}

func (ir *imageResource) RegistryStr() string {
	return ir.ref.Registry
}

func getAuthClient(ctx context.Context, ref registry.Reference, rc registryclient.Client) (*auth.Client, bool, error) {
	if err := rc.RefreshKeychainPullSecrets(ctx); err != nil {
		return nil, false, errors.Wrapf(err, "failed to refresh image pull secrets")
	}

	authn, err := rc.Keychain().Resolve(&imageResource{ref})
	if err != nil {
		return nil, false, errors.Wrapf(err, "failed to resolve auth for %s", ref.String())
	}

	authConfig, err := authn.Authorization()
	if err != nil {
		return nil, false, errors.Wrapf(err, "failed to get auth config for %s", ref.String())
	}

	credentials := auth.Credential{
		Username:     authConfig.Username,
		Password:     authConfig.Password,
		AccessToken:  authConfig.IdentityToken,
		RefreshToken: authConfig.RegistryToken,
	}

	authClient := &auth.Client{
		Credential: func(ctx context.Context, registry string) (auth.Credential, error) {
			switch registry {
			default:
				return credentials, nil
			}
		},
		Cache:    auth.NewCache(),
		ClientID: "notation",
	}

	authClient.SetUserAgent("kyverno.io")
	return authClient, false, nil
}

func getAuthenticator(ctx context.Context, ref string, registryClient registryclient.Client) (*authn.Authenticator, error) {
	parsedRef, err := registry.ParseReference(ref)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse registry reference %s", ref)
	}

	if err := registryClient.RefreshKeychainPullSecrets(ctx); err != nil {
		return nil, errors.Wrapf(err, "failed to refresh image pull secrets")
	}

	authn, err := registryClient.Keychain().Resolve(&imageResource{parsedRef})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve auth for %s", parsedRef.String())
	}
	return &authn, nil
}

func resolveDigest(repo notationregistry.Repository, ref registry.Reference) (registry.Reference, error) {
	if isDigestReference(ref.String()) {
		return ref, nil
	}

	// Resolve tag reference to digest reference.
	manifestDesc, err := getManifestDescriptorFromReference(repo, ref.String())
	if err != nil {
		return registry.Reference{}, err
	}

	ref.Reference = manifestDesc.Digest.String()
	return ref, nil
}

func isDigestReference(reference string) bool {
	parts := strings.SplitN(reference, "/", 2)
	if len(parts) == 1 {
		return false
	}

	index := strings.Index(parts[1], "@")
	return index != -1
}

func getManifestDescriptorFromReference(repo notationregistry.Repository, reference string) (ocispec.Descriptor, error) {
	ref, err := registry.ParseReference(reference)
	if err != nil {
		return ocispec.Descriptor{}, err
	}

	return repo.Resolve(context.Background(), ref.ReferenceOrDefault())
}

func getRemoteOpts(authenticator authn.Authenticator) ([]gcrremote.Option, error) {
	remoteOpts := []gcrremote.Option{}
	remoteOpts = append(remoteOpts, gcrremote.WithAuth(authenticator))

	pusher, err := gcrremote.NewPusher(remoteOpts...)
	if err != nil {
		return nil, err
	}
	remoteOpts = append(remoteOpts, gcrremote.Reuse(pusher))

	puller, err := gcrremote.NewPuller(remoteOpts...)
	if err != nil {
		return nil, err
	}
	remoteOpts = append(remoteOpts, gcrremote.Reuse(puller))

	return remoteOpts, nil
}

func resolveDigestCrane(repo notationregistry.Repository, craneOpts crane.Option, remoteOpts []gcrremote.Option, ref name.Reference) error {
	_, err := repo.Resolve(context.Background(), ref.Name())
	if err != nil {
		return err
	}
	return nil
}
