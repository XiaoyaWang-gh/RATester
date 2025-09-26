package framework

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &request.Request{
		// fill in the request fields
	}

	expect := &RequestExpect{
		req: req,
	}

	resp, err := expect.Do()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp == nil {
		t.Fatalf("Expected response, but got nil")
	}

	// add more assertions as needed
}
