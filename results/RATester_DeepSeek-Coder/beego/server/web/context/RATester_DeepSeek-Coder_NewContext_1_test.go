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
		t.Errorf("Expected a non-nil context, got nil")
	}

	if ctx.Input == nil {
		t.Errorf("Expected a non-nil Input, got nil")
	}

	if ctx.Output == nil {
		t.Errorf("Expected a non-nil Output, got nil")
	}
}
