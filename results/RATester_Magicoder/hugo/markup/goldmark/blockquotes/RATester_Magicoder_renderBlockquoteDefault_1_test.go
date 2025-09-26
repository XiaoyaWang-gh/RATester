package blockquotes

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func TestrenderBlockquoteDefault_1(t *testing.T) {
	tests := []struct {
		name    string
		r       *htmlRenderer
		w       util.BufWriter
		n       ast.Node
		text    string
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

			got, err := tt.r.renderBlockquoteDefault(tt.w, tt.n, tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("htmlRenderer.renderBlockquoteDefault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("htmlRenderer.renderBlockquoteDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
