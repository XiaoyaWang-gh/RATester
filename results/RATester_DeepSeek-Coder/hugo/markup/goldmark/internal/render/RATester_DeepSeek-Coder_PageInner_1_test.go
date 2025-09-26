package render

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	htext "github.com/gohugoio/hugo/common/text"
)

func TestPageInner_1(t *testing.T) {
	type fields struct {
		page            any
		pageInner       any
		ordinal         int
		pos             htext.Position
		posInit         sync.Once
		createPos       func() htext.Position
		getSourceSample func() []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		{
			name: "Test case 1",
			fields: fields{
				page:      "test",
				pageInner: "inner",
			},
			want: "inner",
		},
		{
			name: "Test case 2",
			fields: fields{
				page:      "test",
				pageInner: 123,
			},
			want: 123,
		},
		{
			name: "Test case 3",
			fields: fields{
				page:      "test",
				pageInner: nil,
			},
			want: nil,
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
				page:            tt.fields.page,
				pageInner:       tt.fields.pageInner,
				ordinal:         tt.fields.ordinal,
				pos:             tt.fields.pos,
				posInit:         tt.fields.posInit,
				createPos:       tt.fields.createPos,
				getSourceSample: tt.fields.getSourceSample,
			}
			if got := c.PageInner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hookBase.PageInner() = %v, want %v", got, tt.want)
			}
		})
	}
}
