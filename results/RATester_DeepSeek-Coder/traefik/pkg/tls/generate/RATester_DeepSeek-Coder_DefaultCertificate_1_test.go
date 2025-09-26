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

	t.Parallel()

	cert, err := DefaultCertificate()
	if err != nil {
		t.Fatalf("DefaultCertificate() error = %v", err)
	}

	if cert == nil {
		t.Fatalf("DefaultCertificate() = nil, want not nil")
	}

	if cert.PrivateKey == nil {
		t.Fatalf("DefaultCertificate().PrivateKey = nil, want not nil")
	}

	if len(cert.Certificate) == 0 {
		t.Fatalf("DefaultCertificate().Certificate = empty, want not empty")
	}
}
