package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/types/hstring"
	"github.com/gohugoio/hugo/markup/internal/attributes"
)

func TestPageInner_2(t *testing.T) {
	type fields struct {
		page             any
		pageInner        any
		level            int
		anchor           string
		text             hstring.HTML
		plainText        string
		AttributesHolder *attributes.AttributesHolder
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := headingContext{
				page:             tt.fields.page,
				pageInner:        tt.fields.pageInner,
				level:            tt.fields.level,
				anchor:           tt.fields.anchor,
				text:             tt.fields.text,
				plainText:        tt.fields.plainText,
				AttributesHolder: tt.fields.AttributesHolder,
			}
			if got := ctx.PageInner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("headingContext.PageInner() = %v, want %v", got, tt.want)
			}
		})
	}
}
