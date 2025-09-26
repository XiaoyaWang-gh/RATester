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

	username := "user"
	passwd := "pass"
	expected := "Basic dXNlcjpwYXNz"
	actual := BasicAuth(username, passwd)
	if actual != expected {
		t.Errorf("BasicAuth() = %v, want %v", actual, expected)
	}
}
