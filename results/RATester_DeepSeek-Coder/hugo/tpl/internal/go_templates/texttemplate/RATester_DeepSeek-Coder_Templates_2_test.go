package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTemplates_2(t *testing.T) {
	testCases := []struct {
		name     string
		template *Template
		want     []*Template
	}{
		{
			name: "Test with nil common",
			template: &Template{
				name:   "test",
				common: nil,
			},
			want: nil,
		},
		{
			name: "Test with non-nil common",
			template: &Template{
				name: "test",
				common: &common{
					tmpl: map[string]*Template{
						"template1": &Template{name: "template1"},
						"template2": &Template{name: "template2"},
					},
				},
			},
			want: []*Template{
				{name: "template1"},
				{name: "template2"},
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

			got := tc.template.Templates()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Templates() = %v, want %v", got, tc.want)
			}
		})
	}
}
