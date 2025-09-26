package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestnotAFunction_1(t *testing.T) {
	type args struct {
		args  []parse.Node
		final reflect.Value
	}
	tests := []struct {
		name string
		s    *state
		args args
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

			tt.s.notAFunction(tt.args.args, tt.args.final)
		})
	}
}
