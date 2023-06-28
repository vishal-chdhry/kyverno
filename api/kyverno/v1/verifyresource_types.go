package v1

import (
	"crypto/x509"
	_ "embed"
)

type VerifyResourceOption struct {
	commonOption       `json:""`
	verifyOption       `json:""`
	cosignVerifyOption `json:""`

	SkipObjects ObjectReferenceList `json:"skipObjects,omitempty"`

	Provenance            bool   `json:"-"`
	DisableDryRun         bool   `json:"-"`
	CheckDryRunForApply   bool   `json:"-"`
	CheckMutatingResource bool   `json:"-"`
	DryRunNamespace       string `json:"-"`
}

type verifyOption struct {
	IgnoreFields ObjectFieldBindingList `json:"ignoreFields,omitempty"`
	Signers      SignerList             `json:"signers,omitempty"`

	MaxResourceManifestNum int `json:"maxResourceManifestNum,omitempty"`

	KeyPath               string `json:"-"`
	ResourceBundleRef     string `json:"-"`
	SignatureResourceRef  string `json:"-"`
	ProvenanceResourceRef string `json:"-"`
	UseCache              bool   `json:"-"`
	CacheDir              string `json:"-"`

	annotationKeyToIgnoreFields bool `json:"-"`
}

type commonOption struct {
	AnnotationConfig `json:""`
}

type cosignVerifyOption struct {
	Certificate      string         `json:"-"`
	CertificateChain string         `json:"-"`
	RekorURL         string         `json:"-"`
	OIDCIssuer       string         `json:"-"`
	RootCerts        *x509.CertPool `json:"-"`
	AllowInsecure    bool           `json:"-"`
}

type AnnotationConfig struct {
	AnnotationKeyDomain string `json:"annotationKeyDomain,omitempty"`

	ResourceBundleRefBaseName string `json:"resourceBundleRefBaseName,omitempty"`
	SignatureBaseName         string `json:"signatureBaseName,omitempty"`
	CertificateBaseName       string `json:"certificateBaseName,omitempty"`
	MessageBaseName           string `json:"messageBaseName,omitempty"`
	BundleBaseName            string `json:"bundleBaseName,omitempty"`
}

type ObjectFieldBindingList []ObjectFieldBinding
type SignerList []string
