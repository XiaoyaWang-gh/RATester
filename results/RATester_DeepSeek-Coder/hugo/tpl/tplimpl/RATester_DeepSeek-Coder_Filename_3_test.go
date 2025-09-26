package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestFilename_3(t *testing.T) {
	type fields struct {
		name       string
		template   string
		isText     bool
		isEmbedded bool
		meta       *hugofs.FileMeta
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				name:       "test",
				template:   "test",
				isText:     true,
				isEmbedded: true,
				meta: &hugofs.FileMeta{
					Filename: "test.txt",
				},
			},
			want: "test.txt",
		},
		{
			name: "Test case 2",
			fields: fields{
				name:       "test2",
				template:   "test2",
				isText:     false,
				isEmbedded: false,
				meta: &hugofs.FileMeta{
					Filename: "test2.html",
				},
			},
			want: "test2.html",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tInfo := templateInfo{
				name:       tt.fields.name,
				template:   tt.fields.template,
				isText:     tt.fields.isText,
				isEmbedded: tt.fields.isEmbedded,
				meta:       tt.fields.meta,
			}
			if got := tInfo.Filename(); got != tt.want {
				t.Errorf("templateInfo.Filename() = %v, want %v", got, tt.want)
			}
		})
	}
}
