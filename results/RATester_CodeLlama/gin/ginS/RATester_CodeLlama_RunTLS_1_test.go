package ginS

import (
	"fmt"
	"testing"
)

func TestRunTLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	addr := "127.0.0.1:8080"
	certFile := "cert.pem"
	keyFile := "key.pem"

	// Act
	err := RunTLS(addr, certFile, keyFile)

	// Assert
	if err != nil {
		t.Errorf("RunTLS failed: %v", err)
	}
}
