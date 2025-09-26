package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewTemplateInfo_1(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     templateInfo
	}{
		{
			name:     "test.html",
			template: "<html></html>",
			want: templateInfo{
				name:       "test.html",
				isText:     false,
				isEmbedded: false,
				template:   "<html></html>",
			},
		},
		{
			name:     "embedded/test.html",
			template: "<html></html>",
			want: templateInfo{
				name:       "test.html",
				isText:     false,
				isEmbedded: true,
				template:   "<html></html>",
			},
		},
		{
			name:     "test.txt",
			template: "This is a text template",
			want: templateInfo{
				name:       "test.txt",
				isText:     true,
				isEmbedded: false,
				template:   "This is a text template",
			},
		},
		{
			name:     "embedded/test.txt",
			template: "This is a text template",
			want: templateInfo{
				name:       "test.txt",
				isText:     true,
				isEmbedded: true,
				template:   "This is a text template",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			th := &templateHandler{}
			if got := th.newTemplateInfo(tt.name, tt.template); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTemplateInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
