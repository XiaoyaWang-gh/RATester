package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestevalEmptyInterface_1(t *testing.T) {
	type args struct {
		dot reflect.Value
		n   parse.Node
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
			if got := s.evalEmptyInterface(tt.args.dot, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evalEmptyInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
