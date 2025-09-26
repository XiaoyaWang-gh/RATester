package client

import (
	"fmt"
	"testing"
)

func TestSetCookies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cookie := make(Cookie)
	cookie.SetCookies(map[string]string{
		"key1": "val1",
		"key2": "val2",
	})

	if len(cookie) != 2 {
		t.Errorf("Expected 2 cookies, got %d", len(cookie))
	}

	if cookie["key1"] != "val1" {
		t.Errorf("Expected 'val1', got '%s'", cookie["key1"])
	}

	if cookie["key2"] != "val2" {
		t.Errorf("Expected 'val2', got '%s'", cookie["key2"])
	}
}
