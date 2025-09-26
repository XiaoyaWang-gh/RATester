package client

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestTLSConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		fasthttp: &fasthttp.Client{},
	}

	config := client.TLSConfig()

	if config == nil {
		t.Error("TLSConfig should not be nil")
	}

	if config.MinVersion != tls.VersionTLS12 {
		t.Errorf("Expected MinVersion to be %d, but got %d", tls.VersionTLS12, config.MinVersion)
	}
}
