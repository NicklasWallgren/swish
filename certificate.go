package swish

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
)

func newTLSClientConfig(configuration *Configuration) (*tls.Config, error) {
	caPool, err := createCertPool(configuration.Environment.CertificationFilePath)

	if err != nil {
		return nil, err
	}

	rpCert, err := createCertLeaf(configuration)

	if err != nil {
		return nil, err
	}

	clientCfg := &tls.Config{
		Certificates: []tls.Certificate{*rpCert},
		ClientCAs:    caPool,
		RootCAs:      caPool,
	}

	return clientCfg, nil
}

func createCertPool(certificatePath string) (*x509.CertPool, error) {
	ca, err := ioutil.ReadFile(certificatePath)

	if err != nil {
		return nil, err
	}

	caPool := x509.NewCertPool()
	if caPool.AppendCertsFromPEM(ca) == false {
		return nil, fmt.Errorf("could not append CA Certificate to pool. Invalid certificate")
	}

	return caPool, nil
}

func createCertLeaf(configuration *Configuration) (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(configuration.CertFile, configuration.KeyFile)

	if err != nil {
		return nil, err
	}

	return &cert, nil
}

func getPEMEncodedCertificateData(content []byte, password string) ([]byte, error) {
	blocks, err := pkcs12.ToPEM(content, password)

	if err != nil {
		return nil, err
	}

	var pemData []byte

	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	return pemData, nil
}
