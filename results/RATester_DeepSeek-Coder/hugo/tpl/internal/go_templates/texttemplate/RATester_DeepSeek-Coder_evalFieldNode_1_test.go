package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEvalFieldNode_1(t *testing.T) {
	type args struct {
		dot   reflect.Value
		field *parse.FieldNode
		args  []parse.Node
		final reflect.Value
	}
	tests := []struct {
		name string
		s    *state
		args args
		want reflect.Value
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

			if got := tt.s.evalFieldNode(tt.args.dot, tt.args.field, tt.args.args, tt.args.final); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("state.evalFieldNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
