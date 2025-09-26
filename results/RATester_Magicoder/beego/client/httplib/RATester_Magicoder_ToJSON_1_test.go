package httplib

import (
	"fmt"
	"testing"
)

func TestToJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var v interface{}
	b := &BeegoHTTPRequest{
		url:  "http://example.com",
		body: []byte(`{"key": "value"}`),
	}

	// Act
	err := b.ToJSON(v)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if v.(map[string]interface{})["key"] != "value" {
		t.Errorf("Expected value to be 'value', but got %v", v.(map[string]interface{})["key"])
	}
}
