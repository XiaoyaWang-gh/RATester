package xlog

import (
	"context"
	"fmt"
	"testing"
)

func TestFromContextSafe_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx := context.Background()

	logger := FromContextSafe(ctx)
	if logger == nil {
		t.Fatal("expected logger to not be nil")
	}

	ctxWithLogger := context.WithValue(ctx, xlogKey, logger)
	loggerFromContext := FromContextSafe(ctxWithLogger)
	if loggerFromContext != logger {
		t.Fatal("expected logger from context to be the same as the original logger")
	}
}
