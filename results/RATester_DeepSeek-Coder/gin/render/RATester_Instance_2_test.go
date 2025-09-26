package render

import (
	"fmt"
	"html/template"
	"reflect"
	"testing"
)

func TestInstance_2(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		template *template.Template
		delims   Delims
		input    string
		data     any
		expected Render
	}{
		{
			name:     "success",
			template: template.Must(template.New("test").Parse("Hello, {{.Name}}")),
			delims:   Delims{"{{", "}}"},
			input:    "test",
			data:     struct{ Name string }{"World"},
			expected: HTML{
				Template: template.Must(template.New("test").Parse("Hello, {{.Name}}")),
				Name:     "test",
				Data:     struct{ Name string }{"World"},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		tc := tc // Capture range variable.
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			r := HTMLProduction{
				Template: tc.template,
				Delims:   tc.delims,
			}

			result := r.Instance(tc.input, tc.data)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
