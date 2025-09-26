package accesslog

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRetried_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SaveRetries{}
	req := &http.Request{}
	s.Retried(req, 1)
	if req.Header.Get(RetryAttempts) != "0" {
		t.Errorf("Retried() = %v, want %v", req.Header.Get(RetryAttempts), "0")
	}
	s.Retried(req, 2)
	if req.Header.Get(RetryAttempts) != "1" {
		t.Errorf("Retried() = %v, want %v", req.Header.Get(RetryAttempts), "1")
	}
	s.Retried(req, 3)
	if req.Header.Get(RetryAttempts) != "2" {
		t.Errorf("Retried() = %v, want %v", req.Header.Get(RetryAttempts), "2")
	}
}
