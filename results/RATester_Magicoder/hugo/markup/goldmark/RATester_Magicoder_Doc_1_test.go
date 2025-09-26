package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/tableofcontents"
)

func TestDoc_1(t *testing.T) {
	type fields struct {
		doc any
		toc *tableofcontents.Fragments
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		{
			name: "Test case 1",
			fields: fields{
				doc: "Test Document",
				toc: &tableofcontents.Fragments{},
			},
			want: "Test Document",
		},
		{
			name: "Test case 2",
			fields: fields{
				doc: 123,
				toc: &tableofcontents.Fragments{},
			},
			want: 123,
		},
		{
			name: "Test case 3",
			fields: fields{
				doc: nil,
				toc: &tableofcontents.Fragments{},
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

			p := parserResult{
				doc: tt.fields.doc,
				toc: tt.fields.toc,
			}
			if got := p.Doc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parserResult.Doc() = %v, want %v", got, tt.want)
			}
		})
	}
}
