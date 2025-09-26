package client

import (
	"fmt"
	"testing"
)

func TestSetRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A test response.
	r := &Response{}
	// and: A test request.
	req := &Request{}
	// when: The request is set to the response.
	r.setRequest(req)
	// then: The request is set to the response.
	if r.request != req {
		t.Errorf("The request is not set to the response.")
	}
}
