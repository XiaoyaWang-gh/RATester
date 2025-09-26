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

	c := make(Cookie)
	m := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}
	c.SetCookies(m)
	if len(c) != 2 {
		t.Errorf("SetCookies failed")
	}
}
