package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestAutoLinkURL_1(t *testing.T) {
	type args struct {
		n      *ast.AutoLink
		source []byte
		r      *hookedRenderer
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

			if got := tt.args.r.autoLinkURL(tt.args.n, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("autoLinkURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
