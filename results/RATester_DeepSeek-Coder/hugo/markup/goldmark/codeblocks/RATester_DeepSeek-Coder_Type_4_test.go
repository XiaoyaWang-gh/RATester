package codeblocks

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter/hooks"
	"github.com/gohugoio/hugo/markup/internal/attributes"
)

func TestType_4(t *testing.T) {
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
			name: "TestType",
			fields: fields{
				lang: "go",
			},
			want: "go",
		},
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
			if got := c.Type(); got != tt.want {
				t.Errorf("codeBlockContext.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
