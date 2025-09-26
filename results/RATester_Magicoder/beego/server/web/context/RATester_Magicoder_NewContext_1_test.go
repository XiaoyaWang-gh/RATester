package context

import (
	"fmt"
	"testing"
)

func TestNewContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := NewContext()

	if ctx == nil {
		t.Error("Expected a non-nil context, but got nil")
	}

	if ctx.Input == nil {
		t.Error("Expected a non-nil Input, but got nil")
	}

	if ctx.Output == nil {
		t.Error("Expected a non-nil Output, but got nil")
	}
}
