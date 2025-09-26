package codeblocks

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func TestrenderCodeBlock_2(t *testing.T) {
	type args struct {
		w        util.BufWriter
		src      []byte
		node     ast.Node
		entering bool
	}
	tests := []struct {
		name    string
		args    args
		want    ast.WalkStatus
		wantErr bool
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
			got, err := r.renderCodeBlock(tt.args.w, tt.args.src, tt.args.node, tt.args.entering)
			if (err != nil) != tt.wantErr {
				t.Errorf("htmlRenderer.renderCodeBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("htmlRenderer.renderCodeBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
