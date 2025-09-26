package client

import (
	"fmt"
	"testing"
)

func TestDelCookies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cookie := make(Cookie)
	cookie.SetCookies(map[string]string{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	})

	keys := []string{"key1", "key3"}
	cookie.DelCookies(keys...)

	if len(cookie) != 1 {
		t.Errorf("Expected length of cookie to be 1, got %d", len(cookie))
	}

	if cookie["key2"] != "val2" {
		t.Errorf("Expected value of key2 to be 'val2', got %s", cookie["key2"])
	}
}
