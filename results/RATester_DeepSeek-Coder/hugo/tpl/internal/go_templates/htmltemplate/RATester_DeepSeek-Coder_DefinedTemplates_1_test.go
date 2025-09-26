package template

import (
	"fmt"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestDefinedTemplates_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		template *Template
		want     string
	}{
		{
			name: "success",
			template: &Template{
				text: &template.Template{
					Tree: &parse.Tree{
						Name: "test",
					},
				},
			},
			want: "test",
		},
		{
			name: "failure",
			template: &Template{
				text: &template.Template{
					Tree: &parse.Tree{
						Name: "fail",
					},
				},
			},
			want: "fail",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.template.DefinedTemplates()
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
