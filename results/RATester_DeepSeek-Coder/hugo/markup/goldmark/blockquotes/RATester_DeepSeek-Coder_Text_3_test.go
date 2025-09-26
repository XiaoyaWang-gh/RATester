package blockquotes

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/types/hstring"
)

func TestText_3(t *testing.T) {
	type fields struct {
		text hstring.HTML
	}
	tests := []struct {
		name   string
		fields fields
		want   hstring.HTML
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

			c := &blockquoteContext{
				text: tt.fields.text,
			}
			if got := c.Text(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blockquoteContext.Text() = %v, want %v", got, tt.want)
			}
		})
	}
}
