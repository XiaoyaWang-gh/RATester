package blockquotes

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func TestRenderBlockquote_1(t *testing.T) {
	tests := []struct {
		name     string
		r        *htmlRenderer
		w        util.BufWriter
		src      []byte
		node     ast.Node
		entering bool
		want     ast.WalkStatus
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &htmlRenderer{}
			got, err := r.renderBlockquote(tt.w, tt.src, tt.node, tt.entering)
			if (err != nil) != tt.wantErr {
				t.Errorf("htmlRenderer.renderBlockquote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("htmlRenderer.renderBlockquote() = %v, want %v", got, tt.want)
			}
		})
	}
}
