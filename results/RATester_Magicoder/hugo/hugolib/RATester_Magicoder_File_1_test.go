package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/source"
)

func TestFile_1(t *testing.T) {
	type fields struct {
		f *source.File
	}
	tests := []struct {
		name   string
		fields fields
		want   *source.File
	}{
		{
			name: "Test Case 1",
			fields: fields{
				f: &source.File{},
			},
			want: &source.File{},
		},
		{
			name: "Test Case 2",
			fields: fields{
				f: &source.File{},
			},
			want: &source.File{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &pageMeta{
				f: tt.fields.f,
			}
			if got := p.File(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pageMeta.File() = %v, want %v", got, tt.want)
			}
		})
	}
}
