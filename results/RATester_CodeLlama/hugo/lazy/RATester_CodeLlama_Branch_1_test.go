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

	ini := New()
	ini.Branch(func(ctx context.Context) (any, error) {
		return nil, nil
	})
	if ini.initCount != 1 {
		t.Errorf("expected initCount to be 1, got %d", ini.initCount)
	}
}
