package client

import (
	"fmt"
	"testing"
)

func TestAcquireResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	resp := AcquireResponse()

	if resp == nil {
		t.Fatal("Expected a non-nil response")
	}

	if resp.client != nil {
		t.Errorf("Expected client to be nil, got %v", resp.client)
	}

	if resp.request != nil {
		t.Errorf("Expected request to be nil, got %v", resp.request)
	}

	if resp.RawResponse != nil {
		t.Errorf("Expected RawResponse to be nil, got %v", resp.RawResponse)
	}

	resp.Reset()

	if resp.client != nil {
		t.Errorf("Expected client to be nil after Reset, got %v", resp.client)
	}

	if resp.request != nil {
		t.Errorf("Expected request to be nil after Reset, got %v", resp.request)
	}

	if resp.RawResponse != nil {
		t.Errorf("Expected RawResponse to be nil after Reset, got %v", resp.RawResponse)
	}
}
