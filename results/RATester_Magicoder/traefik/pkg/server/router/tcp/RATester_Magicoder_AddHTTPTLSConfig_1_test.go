package tcp

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestAddHTTPTLSConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := &Router{}

	sniHost := "example.com"
	config := &tls.Config{}

	router.AddHTTPTLSConfig(sniHost, config)

	if router.hostHTTPTLSConfig == nil {
		t.Fatal("hostHTTPTLSConfig should not be nil")
	}

	if _, ok := router.hostHTTPTLSConfig[sniHost]; !ok {
		t.Fatalf("hostHTTPTLSConfig should contain key %s", sniHost)
	}

	if router.hostHTTPTLSConfig[sniHost] != config {
		t.Fatal("hostHTTPTLSConfig should contain the correct config")
	}
}
