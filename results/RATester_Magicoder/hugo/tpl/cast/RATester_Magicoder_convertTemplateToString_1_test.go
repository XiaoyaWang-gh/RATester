package cast

import (
	"fmt"
	"html/template"
	"reflect"
	"testing"
)

func TestconvertTemplateToString_1(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want any
	}{
		{
			name: "Test with template.HTML",
			v:    template.HTML("<html>Hello, World!</html>"),
			want: "<html>Hello, World!</html>",
		},
		{
			name: "Test with template.CSS",
			v:    template.CSS(".class { color: red; }"),
			want: ".class { color: red; }",
		},
		{
			name: "Test with template.HTMLAttr",
			v:    template.HTMLAttr("<a href='#'>Link</a>"),
			want: "<a href='#'>Link</a>",
		},
		{
			name: "Test with template.JS",
			v:    template.JS("console.log('Hello, World!')"),
			want: "console.log('Hello, World!')",
		},
		{
			name: "Test with template.JSStr",
			v:    template.JSStr("'Hello, World!'"),
			want: "'Hello, World!'",
		},
		{
			name: "Test with string",
			v:    "Hello, World!",
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertTemplateToString(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTemplateToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
