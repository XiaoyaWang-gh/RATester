package request

import (
	"fmt"
	"testing"
)

func TestBody_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{
		body: []byte("hello"),
	}
	r.Body([]byte("world"))
	if r.body == nil || string(r.body) != "world" {
		t.Fatal("body not set")
	}
}
