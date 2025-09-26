package client

import (
	"fmt"
	"testing"
)

func TestDel_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := PathParam{}
	key := "key"
	p.Add(key, "val")
	p.Del(key)
	if _, ok := p[key]; ok {
		t.Errorf("Del failed")
	}
}
