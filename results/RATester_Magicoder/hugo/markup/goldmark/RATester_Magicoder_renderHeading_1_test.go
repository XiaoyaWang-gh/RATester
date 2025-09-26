package goldmark

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func TestrenderHeading_1(t *testing.T) {
	type args struct {
		w        util.BufWriter
		source   []byte
		node     ast.Node
		entering bool
	}
	tests := []struct {
		name    string
		r       *hookedRenderer
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

			got, err := tt.r.renderHeading(tt.args.w, tt.args.source, tt.args.node, tt.args.entering)
			if (err != nil) != tt.wantErr {
				t.Errorf("hookedRenderer.renderHeading() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hookedRenderer.renderHeading() = %v, want %v", got, tt.want)
			}
		})
	}
}
