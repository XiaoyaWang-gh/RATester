package lazy

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestBranchWithTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ini := &Init{}

	timeout := 10 * time.Second
	f := func(ctx context.Context) (any, error) {
		// Simulate some work
		time.Sleep(20 * time.Second)
		return "result", nil
	}

	_, err := ini.BranchWithTimeout(timeout, f).Do(ctx)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
