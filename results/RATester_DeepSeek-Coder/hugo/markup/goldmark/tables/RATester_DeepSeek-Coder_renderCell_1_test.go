package tables

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func TestRenderCell_1(t *testing.T) {
	type args struct {
		w        util.BufWriter
		source   []byte
		node     ast.Node
		entering bool
	}
	tests := []struct {
		name    string
		r       *htmlRenderer
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
			got, err := r.renderCell(tt.args.w, tt.args.source, tt.args.node, tt.args.entering)
			if (err != nil) != tt.wantErr {
				t.Errorf("htmlRenderer.renderCell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("htmlRenderer.renderCell() = %v, want %v", got, tt.want)
			}
		})
	}
}
