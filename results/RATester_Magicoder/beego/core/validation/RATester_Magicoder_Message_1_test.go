package validation

import (
	"fmt"
	"testing"
)

func TestMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Result{
		Error: &Error{},
		Ok:    false,
	}

	r.Message("test message")
	if r.Error.Message != "test message" {
		t.Errorf("Expected message to be 'test message', but got '%s'", r.Error.Message)
	}

	r.Message("test %s", "message")
	if r.Error.Message != "test message" {
		t.Errorf("Expected message to be 'test message', but got '%s'", r.Error.Message)
	}
}
