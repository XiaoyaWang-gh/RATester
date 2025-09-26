package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestDefinedTemplates_2(t *testing.T) {
	type fields struct {
		name       string
		Tree       *parse.Tree
		common     *common
		leftDelim  string
		rightDelim string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				name: "test",
				Tree: &parse.Tree{},
				common: &common{
					tmpl: map[string]*Template{
						"template1": &Template{},
						"template2": &Template{},
					},
				},
				leftDelim:  "{{",
				rightDelim: "}}",
			},
			want: "; defined templates are: \"template1\", \"template2\"",
		},
		{
			name: "Test Case 2",
			fields: fields{
				name: "test",
				Tree: &parse.Tree{},
				common: &common{
					tmpl: map[string]*Template{},
				},
				leftDelim:  "{{",
				rightDelim: "}}",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tr := &Template{
				name:       tt.fields.name,
				Tree:       tt.fields.Tree,
				common:     tt.fields.common,
				leftDelim:  tt.fields.leftDelim,
				rightDelim: tt.fields.rightDelim,
			}
			if got := tr.DefinedTemplates(); got != tt.want {
				t.Errorf("DefinedTemplates() = %v, want %v", got, tt.want)
			}
		})
	}
}
