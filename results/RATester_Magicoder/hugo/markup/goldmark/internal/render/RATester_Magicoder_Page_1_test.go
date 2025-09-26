package render

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	htext "github.com/gohugoio/hugo/common/text"
)

func TestPage_1(t *testing.T) {
	type fields struct {
		page      any
		pageInner any
		ordinal   int
		pos       htext.Position
		posInit   sync.Once
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
				page: "test",
			},
			want: "test",
		},
		{
			name: "Test case 2",
			fields: fields{
				page: 123,
			},
			want: 123,
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

			c := &hookBase{
				page:      tt.fields.page,
				pageInner: tt.fields.pageInner,
				ordinal:   tt.fields.ordinal,
				pos:       tt.fields.pos,
				posInit:   tt.fields.posInit,
				createPos: tt.fields.createPos,
			}
			if got := c.Page(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hookBase.Page() = %v, want %v", got, tt.want)
			}
		})
	}
}
