package client

import (
	"fmt"
	"testing"
)

func TestReleaseResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	resp := &Response{
		client: &Client{},
	}
	ReleaseResponse(resp)

	// Assert that the response has been reset
	if resp.client != nil {
		t.Errorf("Expected client to be nil, but got %v", resp.client)
	}
	if resp.request != nil {
		t.Errorf("Expected request to be nil, but got %v", resp.request)
	}
	if resp.RawResponse != nil {
		t.Errorf("Expected RawResponse to be nil, but got %v", resp.RawResponse)
	}
	if resp.cookie != nil {
		t.Errorf("Expected cookie to be nil, but got %v", resp.cookie)
	}
}
