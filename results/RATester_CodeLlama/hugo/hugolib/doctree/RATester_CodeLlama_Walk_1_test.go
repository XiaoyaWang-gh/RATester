package doctree

import (
	"context"
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestWalk_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var (
		ctx = context.Background()
		r   = &NodeShiftTreeWalker[string]{
			Tree: &NodeShiftTree[string]{
				tree: radix.New(),
			},
			Handle: func(s string, v string, exact DimensionFlag) (terminate bool, err error) {
				return false, nil
			},
		}
	)

	if err := r.Walk(ctx); err != nil {
		t.Fatal(err)
	}
}
