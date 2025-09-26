package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func Testjoin_1(t *testing.T) {
	type args struct {
		a        context
		b        context
		node     parse.Node
		nodeName string
	}
	tests := []struct {
		name string
		args args
		want context
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

			if got := join(tt.args.a, tt.args.b, tt.args.node, tt.args.nodeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("join() = %v, want %v", got, tt.want)
			}
		})
	}
}
