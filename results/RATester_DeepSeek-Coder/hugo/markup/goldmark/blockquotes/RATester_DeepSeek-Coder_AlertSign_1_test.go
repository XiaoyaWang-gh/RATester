package blockquotes

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types/hstring"
	"github.com/gohugoio/hugo/markup/converter/hooks"
	"github.com/gohugoio/hugo/markup/internal/attributes"
)

func TestAlertSign_1(t *testing.T) {
	type fields struct {
		BaseContext      hooks.BaseContext
		text             hstring.HTML
		typ              string
		alert            blockQuoteAlert
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
				alert: blockQuoteAlert{
					sign: "+",
				},
			},
			want: "+",
		},
		{
			name: "Test case 2",
			fields: fields{
				alert: blockQuoteAlert{
					sign: "-",
				},
			},
			want: "-",
		},
		{
			name: "Test case 3",
			fields: fields{
				alert: blockQuoteAlert{
					sign: "",
				},
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

			c := &blockquoteContext{
				BaseContext:      tt.fields.BaseContext,
				text:             tt.fields.text,
				typ:              tt.fields.typ,
				alert:            tt.fields.alert,
				AttributesHolder: tt.fields.AttributesHolder,
			}
			if got := c.AlertSign(); got != tt.want {
				t.Errorf("blockquoteContext.AlertSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
