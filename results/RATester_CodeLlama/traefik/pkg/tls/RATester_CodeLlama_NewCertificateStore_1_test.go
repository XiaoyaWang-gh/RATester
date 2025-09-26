package tls

import (
	"fmt"
	"testing"
)

func TestNewCertificateStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	certStore := NewCertificateStore()

	// Assert
	if certStore == nil {
		t.Errorf("Expected a non-nil CertificateStore")
	}

	if certStore.DynamicCerts == nil {
		t.Errorf("Expected a non-nil DynamicCerts")
	}

	if certStore.CertCache == nil {
		t.Errorf("Expected a non-nil CertCache")
	}
}
