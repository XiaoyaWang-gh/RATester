package web

import (
	"fmt"
	"testing"
)

func TestWithResetParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	reset := true
	opt := WithResetParams(reset)
	opts := &filterOpts{}
	opt(opts)
	if opts.resetParams != reset {
		t.Errorf("WithResetParams() failed")
	}
}
