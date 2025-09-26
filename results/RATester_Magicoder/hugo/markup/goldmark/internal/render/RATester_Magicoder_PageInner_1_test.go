package render

import (
	"fmt"
	"reflect"
	"testing"

	htext "github.com/gohugoio/hugo/common/text"
)

func TestPageInner_1(t *testing.T) {
	type fields struct {
		page      any
		pageInner any
		ordinal   int
		pos       htext.Position
		createPos func() htext.Position
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		{
			name: "Test case 1",
			fields: fields{
				page:      "test_page",
				pageInner: "test_page_inner",
				ordinal:   1,
				pos:       htext.Position{},
				createPos: func() htext.Position { return htext.Position{} },
			},
			want: "test_page_inner",
		},
		{
			name: "Test case 2",
			fields: fields{
				page:      "test_page_2",
				pageInner: "test_page_inner_2",
				ordinal:   2,
				pos:       htext.Position{},
				createPos: func() htext.Position { return htext.Position{} },
			},
			want: "test_page_inner_2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &hookBase{
				page:      tt.fields.page,
				pageInner: tt.fields.pageInner,
				ordinal:   tt.fields.ordinal,
				pos:       tt.fields.pos,
				createPos: tt.fields.createPos,
			}
			if got := c.PageInner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hookBase.PageInner() = %v, want %v", got, tt.want)
			}
		})
	}
}
