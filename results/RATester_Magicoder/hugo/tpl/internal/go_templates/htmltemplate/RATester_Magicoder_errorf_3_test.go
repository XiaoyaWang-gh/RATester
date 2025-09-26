package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func Testerrorf_3(t *testing.T) {
	type args struct {
		k    ErrorCode
		node parse.Node
		line int
		f    string
		args []any
	}
	tests := []struct {
		name string
		args args
		want *Error
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

			if got := errorf(tt.args.k, tt.args.node, tt.args.line, tt.args.f, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorf() = %v, want %v", got, tt.want)
			}
		})
	}
}
