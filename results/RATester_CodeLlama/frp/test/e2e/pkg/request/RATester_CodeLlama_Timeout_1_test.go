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

	// given
	timeout := 10 * time.Second
	request := &Request{
		timeout: timeout,
	}

	// when
	request.Timeout(timeout)

	// then
	if request.timeout != timeout {
		t.Errorf("expected %v, actual %v", timeout, request.timeout)
	}
}
