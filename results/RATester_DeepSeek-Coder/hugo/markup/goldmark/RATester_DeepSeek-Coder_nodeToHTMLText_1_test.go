package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestNodeToHTMLText_1(t *testing.T) {
	type args struct {
		n      ast.Node
		source []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
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

			if got := nodeToHTMLText(tt.args.n, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodeToHTMLText() = %v, want %v", got, tt.want)
			}
		})
	}
}
