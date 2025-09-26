package echo

import (
	"fmt"
	"testing"
)

func TestAfter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	fn := func() {
		// Test function
	}
	r.After(fn)
	if len(r.afterFuncs) != 1 {
		t.Errorf("Expected 1 function in afterFuncs, got %d", len(r.afterFuncs))
	}
}
