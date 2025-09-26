package client

import (
	"fmt"
	"testing"
	"time"
)

func TestSetTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	timeout := time.Duration(10)
	r.SetTimeout(timeout)
	if r.timeout != timeout {
		t.Errorf("Expected timeout to be %v, but got %v", timeout, r.timeout)
	}
}
