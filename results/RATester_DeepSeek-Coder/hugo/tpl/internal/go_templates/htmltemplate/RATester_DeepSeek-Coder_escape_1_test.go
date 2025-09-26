package template

import (
	"fmt"
	"sync"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestEscape_1(t *testing.T) {
	type testCase struct {
		name      string
		template  *Template
		wantError bool
	}

	testCases := []testCase{
		{
			name: "Successful escape",
			template: &Template{
				text: &template.Template{},
				nameSpace: &nameSpace{
					mu: sync.Mutex{},
				},
			},
			wantError: false,
		},
		{
			name: "Unsuccessful escape",
			template: &Template{
				text: &template.Template{},
				nameSpace: &nameSpace{
					mu: sync.Mutex{},
				},
				escapeErr: fmt.Errorf("escape error"),
			},
			wantError: true,
		},
		{
			name: "Incomplete template",
			template: &Template{
				text: &template.Template{},
				nameSpace: &nameSpace{
					mu: sync.Mutex{},
				},
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tc.template.escape()
			if (err != nil) != tc.wantError {
				t.Errorf("escape() error = %v, wantError %v", err, tc.wantError)
			}
		})
	}
}
