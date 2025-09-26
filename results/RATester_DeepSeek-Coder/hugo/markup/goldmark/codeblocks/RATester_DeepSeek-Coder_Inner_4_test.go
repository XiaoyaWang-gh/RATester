package codeblocks

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter/hooks"
	"github.com/gohugoio/hugo/markup/internal/attributes"
)

func TestInner_4(t *testing.T) {
	type fields struct {
		BaseContext      hooks.BaseContext
		lang             string
		code             string
		AttributesHolder *attributes.AttributesHolder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				lang: "go",
				code: "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, world\")\n}",
			},
			want: "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, world\")\n}",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &codeBlockContext{
				BaseContext:      tt.fields.BaseContext,
				lang:             tt.fields.lang,
				code:             tt.fields.code,
				AttributesHolder: tt.fields.AttributesHolder,
			}
			if got := c.Inner(); got != tt.want {
				t.Errorf("codeBlockContext.Inner() = %v, want %v", got, tt.want)
			}
		})
	}
}
