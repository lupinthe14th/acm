package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
)

func newTLSConfig(caFile string) (*tls.Config, error) {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()
	if caFile != "" {
		pemCerts, err := os.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		if !certpool.AppendCertsFromPEM(pemCerts) {
			return nil, errors.New("failed to add certs")
		}
	}

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
	}, nil
}
