package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestAt_1(t *testing.T) {
	type args struct {
		node parse.Node
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

			tt.s.at(tt.args.node)
		})
	}
}
