package xlog

import (
	"context"
	"fmt"
	"testing"
)

func TestFromContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx := context.Background()

	// Test with no logger in context
	xl, ok := FromContext(ctx)
	if ok {
		t.Errorf("Expected ok to be false, got true")
	}
	if xl != nil {
		t.Errorf("Expected xl to be nil, got %v", xl)
	}

	// Test with logger in context
	logger := &Logger{}
	ctx = context.WithValue(ctx, xlogKey, logger)
	xl, ok = FromContext(ctx)
	if !ok {
		t.Errorf("Expected ok to be true, got false")
	}
	if xl != logger {
		t.Errorf("Expected xl to be %v, got %v", logger, xl)
	}
}
