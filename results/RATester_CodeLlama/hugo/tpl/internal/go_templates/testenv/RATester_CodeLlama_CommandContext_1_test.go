package testenv

import (
	"context"
	"fmt"
	"testing"
)

func TestCommandContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx := context.Background()
	cmd := CommandContext(t, ctx, "echo", "hello")
	if err := cmd.Run(); err != nil {
		t.Errorf("cmd.Run() failed with %v", err)
	}
}
