package client

import (
	"fmt"
	"testing"
)

func TestSetRootCertificateFromString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}

	// Test with a valid PEM string
	pemString := "-----BEGIN CERTIFICATE-----\nMIIC...\n-----END CERTIFICATE-----"
	client.SetRootCertificateFromString(pemString)

	// Test with an invalid PEM string
	invalidPemString := "invalid-pem-string"
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	client.SetRootCertificateFromString(invalidPemString)
}
