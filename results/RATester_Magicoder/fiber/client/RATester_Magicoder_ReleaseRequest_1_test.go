package client

import (
	"context"
	"fmt"
	"testing"
)

func TestReleaseRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{
		ctx: context.Background(),
		// Initialize other fields as needed
	}

	ReleaseRequest(req)

	// Assert that the request was properly released
	// For example, check if the request is in the pool
	if requestPool.Get() != req {
		t.Error("Request was not properly released")
	}
}
