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
		t.Error("Expected client to be not nil")
	}

	if client.fasthttp == nil {
		t.Error("Expected fasthttp to be not nil")
	}

	if client.header == nil {
		t.Error("Expected header to be not nil")
	}

	if client.params == nil {
		t.Error("Expected params to be not nil")
	}

	if client.cookies == nil {
		t.Error("Expected cookies to be not nil")
	}

	if client.path == nil {
		t.Error("Expected path to be not nil")
	}

	if len(client.userRequestHooks) != 0 {
		t.Error("Expected userRequestHooks to be empty")
	}

	if len(client.builtinRequestHooks) != 3 {
		t.Error("Expected builtinRequestHooks to have 3 elements")
	}

	if len(client.userResponseHooks) != 0 {
		t.Error("Expected userResponseHooks to be empty")
	}

	if len(client.builtinResponseHooks) != 2 {
		t.Error("Expected builtinResponseHooks to have 2 elements")
	}

	if client.jsonMarshal == nil {
		t.Error("Expected jsonMarshal to be not nil")
	}

	if client.jsonUnmarshal == nil {
		t.Error("Expected jsonUnmarshal to be not nil")
	}

	if client.xmlMarshal == nil {
		t.Error("Expected xmlMarshal to be not nil")
	}

	if client.xmlUnmarshal == nil {
		t.Error("Expected xmlUnmarshal to be not nil")
	}

	if client.logger == nil {
		t.Error("Expected logger to be not nil")
	}
}
