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

	ini := &Init{}
	timeout := time.Second
	f := func(ctx context.Context) (any, error) {
		return nil, nil
	}
	ini.AddWithTimeout(timeout, f)
	if ini.initCount != 1 {
		t.Errorf("initCount = %d, want 1", ini.initCount)
	}
}
