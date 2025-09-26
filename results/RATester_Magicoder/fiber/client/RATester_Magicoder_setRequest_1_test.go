package client

import (
	"fmt"
	"testing"
)

func TestsetRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{}
	res := &Response{}
	res.setRequest(req)

	if res.request != req {
		t.Errorf("Expected request to be %v, but got %v", req, res.request)
	}
}
