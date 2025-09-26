package transport

import (
	"crypto/x509"
	"fmt"
	"math/big"
	"testing"
)

func TestNewRandomTLSKeyPair_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	cert := newRandomTLSKeyPair()

	if cert == nil {
		t.Fatal("Expected a certificate, got nil")
	}

	if cert.PrivateKey == nil {
		t.Fatal("Expected a private key, got nil")
	}

	if len(cert.Certificate) == 0 {
		t.Fatal("Expected a certificate chain, got none")
	}

	leafCert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		t.Fatalf("Failed to parse leaf certificate: %v", err)
	}

	if leafCert.SerialNumber.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("Expected serial number 1, got %v", leafCert.SerialNumber)
	}
}
