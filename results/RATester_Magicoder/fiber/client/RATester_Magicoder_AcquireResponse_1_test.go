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

	resp := AcquireResponse()
	defer resp.Close()

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
