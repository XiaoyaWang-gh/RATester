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

	c := make(Cookie)
	key := "key"
	val := "val"
	c.SetCookie(key, val)
	if c[key] != val {
		t.Errorf("SetCookie failed")
	}
}
