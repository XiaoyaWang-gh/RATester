package passthrough

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func TestRenderPassthroughBlock_1(t *testing.T) {
	tests := []struct {
		name     string
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
			got, err := r.renderPassthroughBlock(tt.w, tt.src, tt.node, tt.entering)
			if (err != nil) != tt.wantErr {
				t.Errorf("htmlRenderer.renderPassthroughBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("htmlRenderer.renderPassthroughBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
