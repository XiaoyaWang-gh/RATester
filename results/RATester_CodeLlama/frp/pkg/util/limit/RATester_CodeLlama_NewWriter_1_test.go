package limit

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/time/rate"
)

func TestNewWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &bytes.Buffer{}
	limiter := rate.NewLimiter(rate.Inf, 1)
	writer := NewWriter(w, limiter)
	if writer.w != w {
		t.Errorf("NewWriter() = %v, want %v", writer.w, w)
	}
	if writer.limiter != limiter {
		t.Errorf("NewWriter() = %v, want %v", writer.limiter, limiter)
	}
}
