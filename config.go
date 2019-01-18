package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func ReadCertsBytes() ([]byte, []byte, error) {
	dataCert, err := getCertBytes(flagDataCert, subv.GetString(KeyDataCert))
	if err != nil {
		return nil, nil, err
	}

	caCert, err := getCertBytes(flagCACert, subv.GetString(KeyCACert))
	if err != nil {
		return nil, nil, err
	}

	return dataCert, caCert, nil
}

func getCertBytes(a *string, b string) ([]byte, error) {
	if a != nil && *a != "" {
		return ioutil.ReadFile(*a)
	}
	return []byte(b), nil
}

func NewConfig(tlsInsecure bool) (*sarama.Config, error) {
	data, ca, err := ReadCertsBytes()
	if err != nil {
		return nil, err
	}

	cfg := sarama.NewConfig()
	cfg.Net.TLS.Enable = true
	cfg.Net.TLS.Config, err = NewTLSConfig(data, ca, tlsInsecure)
	return cfg, err
}

func NewConfigCluster(tlsInsecure bool) (*cluster.Config, error) {
	data, ca, err := ReadCertsBytes()
	if err != nil {
		return nil, err
	}

	cfg := cluster.NewConfig()
	cfg.Net.TLS.Enable = true
	cfg.Net.TLS.Config, err = NewTLSConfig(data, ca, tlsInsecure)

	return cfg, err
}

func NewTLSConfig(certBytes, caCertBytes []byte, insecure bool) (*tls.Config, error) {
	cert, err := tls.X509KeyPair(certBytes, certBytes)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCertBytes)
	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: insecure,
	}, nil
}
