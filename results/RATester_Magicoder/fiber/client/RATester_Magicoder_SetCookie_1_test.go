package client

import (
	"fmt"
	"testing"
)

func TestSetCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cookie := make(Cookie)
	key := "testKey"
	val := "testVal"
	cookie.SetCookie(key, val)
	if cookie[key] != val {
		t.Errorf("Expected %s, got %s", val, cookie[key])
	}
}
