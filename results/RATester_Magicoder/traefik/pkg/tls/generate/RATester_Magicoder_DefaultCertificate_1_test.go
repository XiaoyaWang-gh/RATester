package generate

import (
	"fmt"
	"testing"
)

func TestDefaultCertificate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cert, err := DefaultCertificate()
	if err != nil {
		t.Fatalf("DefaultCertificate() failed: %v", err)
	}

	if cert == nil {
		t.Fatalf("DefaultCertificate() returned nil certificate")
	}

	if len(cert.Certificate) == 0 {
		t.Fatalf("DefaultCertificate() returned empty certificate chain")
	}

	if cert.PrivateKey == nil {
		t.Fatalf("DefaultCertificate() returned nil private key")
	}

	// Add more specific tests as needed
}
