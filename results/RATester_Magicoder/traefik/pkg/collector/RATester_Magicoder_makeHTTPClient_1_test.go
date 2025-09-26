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

	if client == nil {
		t.Error("Expected a non-nil client, but got nil")
	}

	if client.Transport == nil {
		t.Error("Expected a non-nil transport, but got nil")
	}

	if client.Transport.(*http.Transport).DialContext == nil {
		t.Error("Expected a non-nil DialContext, but got nil")
	}

	if client.Transport.(*http.Transport).IdleConnTimeout != 90*time.Second {
		t.Errorf("Expected IdleConnTimeout to be 90 seconds, but got %v", client.Transport.(*http.Transport).IdleConnTimeout)
	}

	if client.Transport.(*http.Transport).TLSHandshakeTimeout != 10*time.Second {
		t.Errorf("Expected TLSHandshakeTimeout to be 10 seconds, but got %v", client.Transport.(*http.Transport).TLSHandshakeTimeout)
	}

	if client.Transport.(*http.Transport).ExpectContinueTimeout != 1*time.Second {
		t.Errorf("Expected ExpectContinueTimeout to be 1 second, but got %v", client.Transport.(*http.Transport).ExpectContinueTimeout)
	}
}
