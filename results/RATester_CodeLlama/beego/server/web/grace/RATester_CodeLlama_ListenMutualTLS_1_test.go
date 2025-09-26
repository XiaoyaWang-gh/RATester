package grace

import (
	"fmt"
	"testing"
)

func TestListenMutualTLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	srv := &Server{}
	certFile := "./testdata/cert.pem"
	keyFile := "./testdata/key.pem"
	trustFile := "./testdata/trust.pem"

	// Act
	ln, err := srv.ListenMutualTLS(certFile, keyFile, trustFile)

	// Assert
	if err != nil {
		t.Errorf("ListenMutualTLS() error = %v", err)
		return
	}
	if ln == nil {
		t.Errorf("ListenMutualTLS() ln = %v", ln)
		return
	}
}
