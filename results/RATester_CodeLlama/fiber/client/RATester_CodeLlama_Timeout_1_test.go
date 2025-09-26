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

	r := &Request{
		timeout: 10 * time.Second,
	}
	if r.Timeout() != 10*time.Second {
		t.Errorf("Timeout() = %v, want %v", r.Timeout(), 10*time.Second)
	}
}
