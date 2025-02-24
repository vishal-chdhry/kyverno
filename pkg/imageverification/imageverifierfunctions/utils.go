package imageverifierfunctions

import (
	"github.com/kyverno/kyverno/api/policies.kyverno.io/v1alpha1"
	"github.com/kyverno/kyverno/pkg/imageverification/imagedataloader"
)

func attestorMap(ivpol *v1alpha1.ImageVerificationPolicy) map[string]v1alpha1.Attestor {
	if ivpol == nil {
		return nil
	}

	return arrToMap(ivpol.Spec.Attestors)
}

func attestationMap(ivpol *v1alpha1.ImageVerificationPolicy) map[string]v1alpha1.Attestation {
	if ivpol == nil {
		return nil
	}

	return arrToMap(ivpol.Spec.Attestations)
}

type ARR_TYPE interface {
	GetKey() string
}

func arrToMap[T ARR_TYPE](arr []T) map[string]T {
	m := make(map[string]T)
	for _, v := range arr {
		m[v.GetKey()] = v
	}

	return m
}

func getRemoteOptsFromPolicy(creds *v1alpha1.Credentials) []imagedataloader.Option {
	if creds == nil {
		return nil
	}

	opts := make([]imagedataloader.Option, 0)

	if creds.AllowInsecureRegistry == true {
		opts = append(opts, imagedataloader.WithInsecure(creds.AllowInsecureRegistry))
	}

	if len(creds.Providers) != 0 {
		providers := make([]string, 0, len(creds.Providers))
		for _, v := range providers {
			providers = append(providers, string(v))
		}
		opts = append(opts, imagedataloader.WithCredentialProviders(providers...))
	}

	if len(creds.Secrets) != 0 {
		opts = append(opts, imagedataloader.WithPullSecret(creds.Secrets))
	}

	return opts
}
