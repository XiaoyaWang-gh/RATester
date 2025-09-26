package gin

import (
	"fmt"
	"testing"
)

func TestRunTLS_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()

	certFile := "testdata/cert.pem"
	keyFile := "testdata/key.pem"
	addr := ":8080"

	err := engine.RunTLS(addr, certFile, keyFile)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
