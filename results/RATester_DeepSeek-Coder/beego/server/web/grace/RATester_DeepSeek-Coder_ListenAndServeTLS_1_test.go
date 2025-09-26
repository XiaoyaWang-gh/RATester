package grace

import (
	"fmt"
	"net/http"
	"testing"
)

func TestListenAndServeTLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		Server: &http.Server{},
	}

	certFile := "testdata/cert.pem"
	keyFile := "testdata/key.pem"

	err := srv.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		t.Errorf("ListenAndServeTLS() error = %v", err)
	}
}
