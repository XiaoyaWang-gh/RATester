package client

import (
	"fmt"
	"testing"
)

func TestSetMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	method := "POST"
	r.SetMethod(method)
	if r.method != method {
		t.Errorf("Expected method to be %s, but got %s", method, r.method)
	}
}
