package tls

import (
	"fmt"
	"testing"
)

func TestGetCipherSuites_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ciphers := getCipherSuites()
	if len(ciphers) == 0 {
		t.Errorf("getCipherSuites() returned empty cipher list")
	}
}
