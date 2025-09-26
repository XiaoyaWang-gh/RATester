package tables

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter/hooks"
)

func TestTHead_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &tableContext{
		tHead: []hooks.TableRow{
			// Initialize your test data here
		},
	}

	result := ctx.THead()

	// Add your assertions here
	if len(result) != len(ctx.tHead) {
		t.Errorf("Expected length of result to be %d, but got %d", len(ctx.tHead), len(result))
	}

	for i, row := range result {
		if len(row) != len(ctx.tHead[i]) {
			t.Errorf("Expected length of row %d to be %d, but got %d", i, len(ctx.tHead[i]), len(row))
		}
	}
}
