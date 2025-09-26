package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTemplateExt_1(t *testing.T) {
	tests := []struct {
		name string
		ext  string
		want []string
	}{
		{
			name: "Add .tpl to beeTemplateExt",
			ext:  ".tpl",
			want: []string{".tpl"},
		},
		{
			name: "Add .html to beeTemplateExt",
			ext:  ".html",
			want: []string{".tpl", ".html"},
		},
		{
			name: "Add .tpl to beeTemplateExt again",
			ext:  ".tpl",
			want: []string{".tpl", ".html"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			AddTemplateExt(tt.ext)
			if !reflect.DeepEqual(beeTemplateExt, tt.want) {
				t.Errorf("got %v, want %v", beeTemplateExt, tt.want)
			}
		})
	}
}
