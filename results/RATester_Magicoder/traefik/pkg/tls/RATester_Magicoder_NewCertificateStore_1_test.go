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

	store := NewCertificateStore()

	if store == nil {
		t.Error("Expected a non-nil CertificateStore, got nil")
	}

	if store.DynamicCerts == nil {
		t.Error("Expected a non-nil DynamicCerts, got nil")
	}

	if store.CertCache == nil {
		t.Error("Expected a non-nil CertCache, got nil")
	}
}
