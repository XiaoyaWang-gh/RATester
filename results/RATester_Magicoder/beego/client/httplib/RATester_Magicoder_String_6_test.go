package httplib

import (
	"fmt"
	"testing"
)

func TestString_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{
		url:  "http://example.com",
		body: []byte("test body"),
	}

	expected := "test body"
	result, err := b.String()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
