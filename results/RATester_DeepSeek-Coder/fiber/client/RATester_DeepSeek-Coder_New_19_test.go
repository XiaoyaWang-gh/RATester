package client

import (
	"fmt"
	"testing"
)

func TestNew_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := New()

	if client == nil {
		t.Errorf("Expected a non-nil client, got nil")
	}

	if client.fasthttp == nil {
		t.Errorf("Expected a non-nil fasthttp client, got nil")
	}

	if client.header == nil {
		t.Errorf("Expected a non-nil header, got nil")
	}

	if client.params == nil {
		t.Errorf("Expected a non-nil params, got nil")
	}

	if client.cookies == nil {
		t.Errorf("Expected a non-nil cookies, got nil")
	}

	if client.path == nil {
		t.Errorf("Expected a non-nil path, got nil")
	}

	if client.jsonMarshal == nil {
		t.Errorf("Expected a non-nil jsonMarshal, got nil")
	}

	if client.jsonUnmarshal == nil {
		t.Errorf("Expected a non-nil jsonUnmarshal, got nil")
	}

	if client.xmlMarshal == nil {
		t.Errorf("Expected a non-nil xmlMarshal, got nil")
	}

	if client.xmlUnmarshal == nil {
		t.Errorf("Expected a non-nil xmlUnmarshal, got nil")
	}

	if client.logger == nil {
		t.Errorf("Expected a non-nil logger, got nil")
	}

	if len(client.userRequestHooks) != 0 {
		t.Errorf("Expected an empty userRequestHooks, got %v", client.userRequestHooks)
	}

	if len(client.builtinRequestHooks) != 3 {
		t.Errorf("Expected 3 builtinRequestHooks, got %v", len(client.builtinRequestHooks))
	}

	if len(client.userResponseHooks) != 0 {
		t.Errorf("Expected an empty userResponseHooks, got %v", client.userResponseHooks)
	}

	if len(client.builtinResponseHooks) != 2 {
		t.Errorf("Expected 2 builtinResponseHooks, got %v", len(client.builtinResponseHooks))
	}
}
