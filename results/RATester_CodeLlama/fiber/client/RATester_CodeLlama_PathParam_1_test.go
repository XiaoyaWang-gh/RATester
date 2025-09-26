package client

import (
	"fmt"
	"testing"
)

func TestPathParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.path = &PathParam{
		"key": "value",
	}

	if r.PathParam("key") != "value" {
		t.Errorf("PathParam() = %v, want %v", r.PathParam("key"), "value")
	}
}
