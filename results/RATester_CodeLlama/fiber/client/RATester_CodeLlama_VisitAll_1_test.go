package client

import (
	"fmt"
	"testing"
)

func TestVisitAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := PathParam{}
	p.Add("key1", "val1")
	p.Add("key2", "val2")
	p.Add("key3", "val3")
	p.VisitAll(func(key string, val string) {
		t.Logf("key: %s, val: %s", key, val)
	})
}
