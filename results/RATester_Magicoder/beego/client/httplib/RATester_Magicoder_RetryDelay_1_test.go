package httplib

import (
	"fmt"
	"testing"
	"time"
)

func TestRetryDelay_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &BeegoHTTPRequest{
		setting: BeegoHTTPSettings{
			RetryDelay: 0,
		},
	}

	delay := time.Second * 5
	req.RetryDelay(delay)

	if req.setting.RetryDelay != delay {
		t.Errorf("Expected RetryDelay to be %v, but got %v", delay, req.setting.RetryDelay)
	}
}
