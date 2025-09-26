package client

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{
		timeout: 10 * time.Second,
	}

	if req.Timeout() != 10*time.Second {
		t.Errorf("Expected timeout to be 10 seconds, but got %v", req.Timeout())
	}
}
