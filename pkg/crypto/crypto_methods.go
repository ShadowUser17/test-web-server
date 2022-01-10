package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	_ "encoding/pem"
	"fmt"
	"math/big"
	"time"
)

func MakeRootCertificate() (*x509.Certificate, *rsa.PrivateKey, error) {
	var now = time.Now()

	var ca_cert = x509.Certificate{
		Subject: pkix.Name{
			Organization:  []string{"-"},
			Country:       []string{"UA"},
			Province:      []string{"-"},
			Locality:      []string{"-"},
			StreetAddress: []string{"-"},
			PostalCode:    []string{"-"},
		},
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		SerialNumber:          big.NewInt(now.Unix()),
		NotBefore:             now,
		NotAfter:              now.AddDate(1, 0, 0),
		IsCA:                  true,
		BasicConstraintsValid: true,
	}

	var cert_key *rsa.PrivateKey
	var err error

	if cert_key, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
		return nil, nil, fmt.Errorf("rsa.GenerateKey: %v", err)
	}

	return &ca_cert, cert_key, nil
}
