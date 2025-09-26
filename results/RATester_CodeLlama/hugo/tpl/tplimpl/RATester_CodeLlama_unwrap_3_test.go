package tplimpl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl"
	template "github.com/gohugoio/hugo/tpl/internal/go_templates/htmltemplate"
)

func TestUnwrap_3(t *testing.T) {
	type args struct {
		templ tpl.Template
	}
	tests := []struct {
		name string
		args args
		want tpl.Template
	}{
		{
			name: "unwrap",
			args: args{
				templ: &templateState{
					Template: &template.Template{},
				},
			},
			want: &template.Template{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := unwrap(tt.args.templ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
