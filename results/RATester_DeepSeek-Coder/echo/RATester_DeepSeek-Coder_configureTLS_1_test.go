package echo

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestConfigureTLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{
		TLSServer: &http.Server{
			TLSConfig: &tls.Config{
				NextProtos: []string{"h2"},
			},
		},
	}

	address := ":1234"
	e.configureTLS(address)

	if e.TLSServer.Addr != address {
		t.Errorf("Expected address to be %s, got %s", address, e.TLSServer.Addr)
	}

	if !strings.Contains(e.TLSServer.TLSConfig.NextProtos[0], "h2") {
		t.Errorf("Expected NextProtos to contain 'h2', got %s", e.TLSServer.TLSConfig.NextProtos[0])
	}
}
