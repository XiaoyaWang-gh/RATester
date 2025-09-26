package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl"
	htmltemplate "github.com/gohugoio/hugo/tpl/internal/go_templates/htmltemplate"
	texttemplate "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestIsText_2(t *testing.T) {
	type args struct {
		templ tpl.Template
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_case_1",
			args: args{
				templ: &texttemplate.Template{},
			},
			want: true,
		},
		{
			name: "test_case_2",
			args: args{
				templ: &htmltemplate.Template{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isText(tt.args.templ); got != tt.want {
				t.Errorf("isText() = %v, want %v", got, tt.want)
			}
		})
	}
}
