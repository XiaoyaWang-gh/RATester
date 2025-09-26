package goldmark

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types/hstring"
	"github.com/gohugoio/hugo/markup/internal/attributes"
)

func TestDestination_1(t *testing.T) {
	type fields struct {
		page             any
		pageInner        any
		destination      string
		title            string
		text             hstring.HTML
		plainText        string
		AttributesHolder *attributes.AttributesHolder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Destination",
			fields: fields{
				destination: "https://example.com",
			},
			want: "https://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := linkContext{
				page:             tt.fields.page,
				pageInner:        tt.fields.pageInner,
				destination:      tt.fields.destination,
				title:            tt.fields.title,
				text:             tt.fields.text,
				plainText:        tt.fields.plainText,
				AttributesHolder: tt.fields.AttributesHolder,
			}
			if got := ctx.Destination(); got != tt.want {
				t.Errorf("linkContext.Destination() = %v, want %v", got, tt.want)
			}
		})
	}
}
