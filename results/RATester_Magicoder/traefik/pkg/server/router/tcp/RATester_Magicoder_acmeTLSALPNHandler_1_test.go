package tcp

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestAcmeTLSALPNHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := &Router{
		httpsTLSConfig: &tls.Config{},
	}

	handler := router.acmeTLSALPNHandler()

	if handler == nil {
		t.Error("Expected a non-nil handler, but got nil")
	}
}
