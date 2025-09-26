package ratelimit

import (
	"fmt"
	"testing"
)

func TestWithRejectionResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	resp := RejectionResponse{
		code: 403,
		body: "rejected",
	}
	opt := WithRejectionResponse(resp)
	l := &limiter{}
	opt(l)
	if l.resp.code != resp.code {
		t.Errorf("Expected %d, got %d", resp.code, l.resp.code)
	}
	if l.resp.body != resp.body {
		t.Errorf("Expected %s, got %s", resp.body, l.resp.body)
	}
}
