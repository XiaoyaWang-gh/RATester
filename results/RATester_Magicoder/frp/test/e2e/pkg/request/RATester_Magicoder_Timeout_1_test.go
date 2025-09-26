package request

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

	r := &Request{}
	timeout := 10 * time.Second
	r.Timeout(timeout)
	if r.timeout != timeout {
		t.Errorf("Expected timeout to be %v, but got %v", timeout, r.timeout)
	}
}
