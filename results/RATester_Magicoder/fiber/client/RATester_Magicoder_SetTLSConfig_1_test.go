package client

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestSetTLSConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	config := &tls.Config{}
	client.SetTLSConfig(config)

	if client.fasthttp.TLSConfig != config {
		t.Errorf("Expected TLSConfig to be set, but it was not.")
	}
}
