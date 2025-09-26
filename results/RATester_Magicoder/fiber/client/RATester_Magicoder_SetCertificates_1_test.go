package client

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestSetCertificates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}

	// Create a test certificate
	cert := &tls.Certificate{
		Certificate: [][]byte{[]byte("test cert")},
		PrivateKey:  []byte("test key"),
	}

	// Call the method under test
	client.SetCertificates(*cert)

	// Assert that the certificate was added to the client
	if len(client.TLSConfig().Certificates) != 1 {
		t.Errorf("Expected 1 certificate, got %d", len(client.TLSConfig().Certificates))
	}
}
