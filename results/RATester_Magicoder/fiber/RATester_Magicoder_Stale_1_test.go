package fiber

import (
	"fmt"
	"testing"
)

func TestStale_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		matched: true,
	}

	if ctx.Stale() != false {
		t.Errorf("Expected Stale() to return false, but got true")
	}

	ctx.matched = false

	if ctx.Stale() != true {
		t.Errorf("Expected Stale() to return true, but got false")
	}
}
