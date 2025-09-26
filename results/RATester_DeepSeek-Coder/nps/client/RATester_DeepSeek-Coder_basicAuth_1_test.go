package client

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
	password := "testpass"
	expected := "dGVzdHVzZXJAZ21haWwuY29tOnRlc3R1c2VycGFzc3dvcmQ="

	result := basicAuth(username, password)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
