package client

import (
	"fmt"
	"testing"
)

func TestResetPathParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{
		path: &PathParam{},
	}

	req.path.Add("key", "value")

	req.ResetPathParams()

	if len(*req.path) != 0 {
		t.Errorf("Expected path params to be reset, but got: %v", *req.path)
	}
}
