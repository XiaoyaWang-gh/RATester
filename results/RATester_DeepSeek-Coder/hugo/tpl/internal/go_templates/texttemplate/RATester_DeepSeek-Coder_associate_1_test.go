package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestAssociate_1(t *testing.T) {
	type args struct {
		new  *Template
		tree *parse.Tree
	}
	tests := []struct {
		name string
		t    *Template
		args args
		want bool
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

			if got := tt.t.associate(tt.args.new, tt.args.tree); got != tt.want {
				t.Errorf("Template.associate() = %v, want %v", got, tt.want)
			}
		})
	}
}
