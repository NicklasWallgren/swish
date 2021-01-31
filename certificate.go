package swish

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/pkcs12"
)

func newTLSClientConfig(configuration *Configuration) (*tls.Config, error) {
	caPool, err := createCertPool(configuration.Environment.Certificate)
	if err != nil {
		return nil, err
	}

	rpCert, err := createCertLeaf(configuration)
	if err != nil {
		return nil, err
	}

	// #nosec:G402
	clientCfg := &tls.Config{
		Certificates: []tls.Certificate{*rpCert},
		ClientCAs:    caPool,
		RootCAs:      caPool,
	}

	return clientCfg, nil
}

func createCertPool(base64EncodedCertificate string) (*x509.CertPool, error) {
	certificate, err := base64.StdEncoding.DecodeString(base64EncodedCertificate)
	if err != nil {
		return nil, fmt.Errorf("could not decode the certificate. %w", err)
	}

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(certificate) {
		return nil, fmt.Errorf("could not append CA Certificate to pool. Invalid base64EncodedCertificate")
	}

	return caPool, nil
}

func createCertLeaf(configuration *Configuration) (*tls.Certificate, error) {
	blocks, err := pkcs12.ToPEM(configuration.Pkcs12.Content, configuration.Pkcs12.Password)
	if err != nil {
		return nil, fmt.Errorf("unable to load pkcs12 %w", err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return nil, fmt.Errorf("unable to load pkcs12 %w", err)
	}

	return &cert, nil
}
