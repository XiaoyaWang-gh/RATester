package types

import (
	"fmt"
	"testing"
)

func TestLoadKeyPair_1(t *testing.T) {
	t.Run("should return error when cert and key are not provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := loadKeyPair("", "")
		if err == nil {
			t.Errorf("Expected error when cert and key are not provided")
		}
	})

	t.Run("should return error when cert file does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := loadKeyPair("/non-existing-cert-file", "key")
		if err == nil {
			t.Errorf("Expected error when cert file does not exist")
		}
	})

	t.Run("should return error when key file does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := loadKeyPair("cert", "/non-existing-key-file")
		if err == nil {
			t.Errorf("Expected error when key file does not exist")
		}
	})

	t.Run("should return key pair when both cert and key files exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		keyPair, err := loadKeyPair("testdata/cert.pem", "testdata/key.pem")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if keyPair.Certificate == nil {
			t.Errorf("Expected certificate, got nil")
		}
		if keyPair.PrivateKey == nil {
			t.Errorf("Expected private key, got nil")
		}
	})
}
