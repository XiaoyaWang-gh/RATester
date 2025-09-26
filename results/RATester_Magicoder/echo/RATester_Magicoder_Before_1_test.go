package echo

import (
	"fmt"
	"testing"
)

func TestBefore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	fn := func() {
		// Test function
	}
	r.Before(fn)
	if len(r.beforeFuncs) != 1 {
		t.Errorf("Expected 1 function in beforeFuncs, got %d", len(r.beforeFuncs))
	}
}
