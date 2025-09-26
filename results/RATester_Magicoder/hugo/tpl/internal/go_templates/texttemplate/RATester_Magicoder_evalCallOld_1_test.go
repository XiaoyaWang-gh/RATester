package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestevalCallOld_1(t *testing.T) {
	type args struct {
		dot       reflect.Value
		fun       reflect.Value
		isBuiltin bool
		node      parse.Node
		name      string
		args      []parse.Node
		final     reflect.Value
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
			if got := s.evalCallOld(tt.args.dot, tt.args.fun, tt.args.isBuiltin, tt.args.node, tt.args.name, tt.args.args, tt.args.final); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evalCallOld() = %v, want %v", got, tt.want)
			}
		})
	}
}
