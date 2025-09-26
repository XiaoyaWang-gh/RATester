package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestIdentifierBase_3(t *testing.T) {
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
				name: "test",
			},
			want: "test",
		},
		{
			name: "Test case 2",
			fields: fields{
				name: "test2",
			},
			want: "test2",
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
			if got := tInfo.IdentifierBase(); got != tt.want {
				t.Errorf("templateInfo.IdentifierBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
