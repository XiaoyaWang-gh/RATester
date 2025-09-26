package grace

import (
	"fmt"
	"net/http"
	"testing"
)

func TestListenAndServeMutualTLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		Server: &http.Server{},
	}

	certFile := "testdata/server.crt"
	keyFile := "testdata/server.key"
	trustFile := "testdata/ca.crt"

	err := srv.ListenAndServeMutualTLS(certFile, keyFile, trustFile)
	if err != nil {
		t.Errorf("ListenAndServeMutualTLS() error = %v", err)
	}
}
