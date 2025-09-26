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

	cookie.DelCookies("key1", "key2")

	if len(cookie) != 1 {
		t.Errorf("Expected length of cookie to be 1, but got %d", len(cookie))
	}

	if _, ok := cookie["key1"]; ok {
		t.Error("Expected key1 to be deleted, but it still exists")
	}

	if _, ok := cookie["key2"]; ok {
		t.Error("Expected key2 to be deleted, but it still exists")
	}

	if _, ok := cookie["key3"]; !ok {
		t.Error("Expected key3 to still exist, but it was deleted")
	}
}
