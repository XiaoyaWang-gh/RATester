package template

import (
	"fmt"
	"reflect"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestPrepare_1(t *testing.T) {
	tests := []struct {
		name string
		t    *Template
		want *template.Template
	}{
		{
			name: "ok",
			t: &Template{
				text: &template.Template{},
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

			if got, _ := tt.t.Prepare(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.Prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}
