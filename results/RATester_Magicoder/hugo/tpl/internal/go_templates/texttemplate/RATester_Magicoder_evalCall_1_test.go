package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestevalCall_1(t *testing.T) {
	type args struct {
		dot       reflect.Value
		fun       reflect.Value
		isBuiltin bool
		node      parse.Node
		name      string
		args      []parse.Node
		final     reflect.Value
		first     []reflect.Value
	}
	tests := []struct {
		name string
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

			s := &state{}
			if got := s.evalCall(tt.args.dot, tt.args.fun, tt.args.isBuiltin, tt.args.node, tt.args.name, tt.args.args, tt.args.final, tt.args.first...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evalCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
