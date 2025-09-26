package request

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestTLSConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Prepare
	var (
		tlsConfig = &tls.Config{}
		request   = &Request{}
	)

	// Action
	request.TLSConfig(tlsConfig)

	// Assert
	if request.tlsConfig != tlsConfig {
		t.Errorf("request.tlsConfig = %v, want %v", request.tlsConfig, tlsConfig)
	}
}
