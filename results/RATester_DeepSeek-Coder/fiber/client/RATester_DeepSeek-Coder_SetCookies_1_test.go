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
		t.Errorf("Expected length of cookie to be 2, got %d", len(cookie))
	}

	if cookie["key1"] != "val1" {
		t.Errorf("Expected value of key1 to be val1, got %s", cookie["key1"])
	}

	if cookie["key2"] != "val2" {
		t.Errorf("Expected value of key2 to be val2, got %s", cookie["key2"])
	}
}
