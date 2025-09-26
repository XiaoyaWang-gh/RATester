package template

import (
	"fmt"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestName_5(t *testing.T) {
	type fields struct {
		escapeErr error
		text      *template.Template
		Tree      *parse.Tree
		nameSpace *nameSpace
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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

			tmpl := &Template{
				escapeErr: tt.fields.escapeErr,
				text:      tt.fields.text,
				Tree:      tt.fields.Tree,
				nameSpace: tt.fields.nameSpace,
			}
			if got := tmpl.Name(); got != tt.want {
				t.Errorf("Template.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
