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

	c := make(Cookie)
	c.Add("key1", "val1")
	c.Add("key2", "val2")
	c.Add("key3", "val3")
	c.DelCookies("key1", "key2")
	if len(c) != 1 {
		t.Errorf("DelCookies failed")
	}
}
