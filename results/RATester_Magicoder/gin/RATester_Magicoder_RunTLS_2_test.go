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

	engine := &Engine{}

	addr := ":8080"
	certFile := "testdata/server.pem"
	keyFile := "testdata/server.key"

	err := engine.RunTLS(addr, certFile, keyFile)
	if err != nil {
		t.Errorf("RunTLS returned an error: %v", err)
	}
}
