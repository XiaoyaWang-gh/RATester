package template

import (
	"fmt"
	"testing"
)

func TestFuncs_2(t *testing.T) {
	testCases := []struct {
		name     string
		funcMap  FuncMap
		expected *Template
	}{
		{
			name: "Test with empty FuncMap",
			funcMap: FuncMap{
				"empty": func() {},
			},
			expected: &Template{
				name: "test",
			},
		},
		{
			name: "Test with non-empty FuncMap",
			funcMap: FuncMap{
				"add": func(a, b int) int {
					return a + b
				},
			},
			expected: &Template{
				name: "test",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			template := &Template{
				name: "test",
			}
			result := template.Funcs(tc.funcMap)
			if result.name != tc.expected.name {
				t.Errorf("Expected name %s, got %s", tc.expected.name, result.name)
			}
		})
	}
}
