package collector

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestMakeHTTPClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := makeHTTPClient()

	if client.Transport == nil {
		t.Errorf("Expected a non-nil Transport, got nil")
	}

	if client.Timeout != 30*time.Second {
		t.Errorf("Expected a timeout of 30 seconds, got %v", client.Timeout)
	}

	if client.Transport.(*http.Transport).IdleConnTimeout != 90*time.Second {
		t.Errorf("Expected an IdleConnTimeout of 90 seconds, got %v", client.Transport.(*http.Transport).IdleConnTimeout)
	}

	if client.Transport.(*http.Transport).TLSHandshakeTimeout != 10*time.Second {
		t.Errorf("Expected a TLSHandshakeTimeout of 10 seconds, got %v", client.Transport.(*http.Transport).TLSHandshakeTimeout)
	}

	if client.Transport.(*http.Transport).ExpectContinueTimeout != 1*time.Second {
		t.Errorf("Expected an ExpectContinueTimeout of 1 second, got %v", client.Transport.(*http.Transport).ExpectContinueTimeout)
	}
}
