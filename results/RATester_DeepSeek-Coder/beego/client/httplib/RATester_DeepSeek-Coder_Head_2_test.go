package httplib

import (
	"fmt"
	"testing"
)

func TestHead_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		Name:     "test",
		Endpoint: "http://localhost:8080",
	}

	type testStruct struct {
		Name string `json:"name"`
	}

	var result testStruct
	err := client.Head(&result, "/test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.Name != "test" {
		t.Errorf("Expected name to be 'test', got %s", result.Name)
	}
}
