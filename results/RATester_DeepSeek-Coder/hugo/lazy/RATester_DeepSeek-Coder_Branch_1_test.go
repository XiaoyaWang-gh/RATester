package lazy

import (
	"context"
	"fmt"
	"testing"
)

func TestBranch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx := context.Background()
	init := New()

	initFn := func(ctx context.Context) (any, error) {
		return "test", nil
	}

	init.Branch(initFn)

	out, err := init.Do(ctx)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if out != "test" {
		t.Errorf("expected 'test', got %v", out)
	}
}
