package client

import (
	"fmt"
	"testing"
)

func TestbasicAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	username := "testuser"
	password := "testpassword"
	expected := "dGVzdHVzZXI6dGVzdHBhc3N3b3Jk"

	result := basicAuth(username, password)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
