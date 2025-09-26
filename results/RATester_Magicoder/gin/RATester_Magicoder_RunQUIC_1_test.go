package gin

import (
	"fmt"
	"testing"
)

func TestRunQUIC_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{}

	addr := ":8080"
	certFile := "cert.pem"
	keyFile := "key.pem"

	err := engine.RunQUIC(addr, certFile, keyFile)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
