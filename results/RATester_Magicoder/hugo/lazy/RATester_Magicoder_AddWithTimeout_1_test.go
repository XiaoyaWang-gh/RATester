package lazy

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestAddWithTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ini := &Init{}

	timeout := 10 * time.Second
	f := func(ctx context.Context) (any, error) {
		return "test", nil
	}

	result, err := ini.AddWithTimeout(timeout, f).Do(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != "test" {
		t.Errorf("Expected 'test', got %v", result)
	}
}
