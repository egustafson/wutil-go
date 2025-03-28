package wlib

import (
	"crypto/tls"

	_ "github.com/go-viper/mapstructure/v2"
)

// TlsProfile holds the raw input used to build a TlsConfig, (which is an
// extension of a crypto/tls.Config).
type TlsProfile struct {
	// Cert holds either the path to, or the complete PEM data of the
	// certificate, including all necessary intermediaries.  The leaf
	// certificate must be the first PEM block in either the file or the data.
	Cert string `mapstructure:"cert,omitempty"`
	// Key holds either the path to, or PEM data for the private key paired with
	// `Cert`.
	Key string `mapstructure:"key,omitempty"`
	// CA holds either the path to, or the complete PEM data of certificates to
	// be considered authoritative signers of certificates.  If the CA field is
	// an empty string then the ClientCAs and RootCAs fields in tls.Config will
	// be null.
	CA string `mapstructure:"ca,omitempty"`
	// DisableTLS is a flag that indicates TLS would not be used, it is mostly
	// informational.  If DisableTls is set then the Cert, Key, and CA fields
	// may be empty and they will not be processed by functions in this module.
	DisableTLS bool `mapstructure:"disable-tls,omitempty"`
	// DisableValidation is a flag that indicates that connections with remote
	// agents should not be validated against the CA and the CA field may be
	// blank.
	DisableValidation bool `mapstructure:"disable-validation,omitempty"`

	tlsConfig *tls.Config
}

func (tlsProfile *TlsProfile) TlsConfig() *tls.Config {
	//
	// TODO: implement
	//
	return tlsProfile.tlsConfig
}
