package httplib

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestWithTLSClientConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &tls.Config{
		// configure tls.Config here
	}
	client := &Client{}
	WithTLSClientConfig(config)(client)

	// assert that the client's TLSClientConfig is set correctly
	if client.Setting.TLSClientConfig != config {
		t.Errorf("Expected TLSClientConfig to be set correctly, but it was not.")
	}
}
