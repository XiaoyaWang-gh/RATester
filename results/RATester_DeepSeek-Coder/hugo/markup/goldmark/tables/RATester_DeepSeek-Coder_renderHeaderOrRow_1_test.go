package tables

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter/hooks"
	"github.com/gohugoio/hugo/markup/goldmark/internal/render"
	"github.com/yuin/goldmark/ast"
)

func TestRenderHeaderOrRow_1(t *testing.T) {
	testCases := []struct {
		name     string
		source   []byte
		expected []hooks.TableRow
	}{
		{
			name:     "Test with table header",
			source:   []byte("## Header"),
			expected: []hooks.TableRow{{}},
		},
		{
			name:     "Test with table row",
			source:   []byte("Row content"),
			expected: []hooks.TableRow{{}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &htmlRenderer{}
			ctx := &render.Context{}
			node := &ast.Heading{Level: 2}

			_, err := r.renderHeaderOrRow(ctx, tc.source, node, true)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			table := r.peekTable(ctx)
			if len(table.THead) != len(tc.expected) {
				t.Errorf("Expected %d table rows, got %d", len(tc.expected), len(table.THead))
			}

			_, err = r.renderHeaderOrRow(ctx, tc.source, node, false)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			table = r.peekTable(ctx)
			if len(table.TBody) != len(tc.expected) {
				t.Errorf("Expected %d table rows, got %d", len(tc.expected), len(table.TBody))
			}
		})
	}
}
