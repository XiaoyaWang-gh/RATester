package transport

import (
	"fmt"
	"testing"
)

func TestNewRandomTLSKeyPair_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cert := newRandomTLSKeyPair()
	if cert == nil {
		t.Error("newRandomTLSKeyPair() returned nil")
	}
	if cert.PrivateKey == nil {
		t.Error("newRandomTLSKeyPair() returned certificate with nil private key")
	}
	if len(cert.Certificate) == 0 {
		t.Error("newRandomTLSKeyPair() returned certificate with no certificates")
	}
}
