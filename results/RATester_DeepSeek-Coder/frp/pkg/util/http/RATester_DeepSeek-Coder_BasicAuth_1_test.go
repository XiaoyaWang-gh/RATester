package http

import (
	"fmt"
	"testing"
)

func TestBasicAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	username := "testuser"
	passwd := "testpass"
	expected := "Basic dGVzdHVzZXJAZ21haWwuY29tOjEyMzEyMw=="
	result := BasicAuth(username, passwd)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
