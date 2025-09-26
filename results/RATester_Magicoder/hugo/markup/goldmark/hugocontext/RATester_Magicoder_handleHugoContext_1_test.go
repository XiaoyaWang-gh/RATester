package hugocontext

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func TesthandleHugoContext_1(t *testing.T) {
	type args struct {
		w        util.BufWriter
		source   []byte
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

			r := &hugoContextRenderer{}
			got, err := r.handleHugoContext(tt.args.w, tt.args.source, tt.args.node, tt.args.entering)
			if (err != nil) != tt.wantErr {
				t.Errorf("hugoContextRenderer.handleHugoContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hugoContextRenderer.handleHugoContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
