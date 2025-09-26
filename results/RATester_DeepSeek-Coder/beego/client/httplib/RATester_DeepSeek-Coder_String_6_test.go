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
		url: "http://example.com",
	}

	expected := "http://example.com"
	result, err := b.String()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
