package httplib

import (
	"fmt"
	"testing"
)

func TestPut_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		Name:     "test",
		Endpoint: "http://localhost:8080",
	}

	value := struct{}{}
	path := "/test"
	body := struct{}{}
	opts := []BeegoHTTPRequestOption{}

	err := client.Put(value, path, body, opts...)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
