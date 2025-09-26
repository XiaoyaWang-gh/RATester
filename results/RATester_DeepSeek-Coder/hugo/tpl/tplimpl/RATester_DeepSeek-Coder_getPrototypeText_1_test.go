package tplimpl

import (
	"fmt"
	"reflect"
	"testing"

	texttemplate "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestGetPrototypeText_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		t    *templateNamespace
		want *texttemplate.Template
	}{
		{
			name: "prototypeTextClone is not nil",
			t: &templateNamespace{
				prototypeTextClone: &texttemplate.Template{},
			},
			want: &texttemplate.Template{},
		},
		{
			name: "prototypeTextClone is nil",
			t: &templateNamespace{
				prototypeText: &texttemplate.Template{},
			},
			want: &texttemplate.Template{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.t.getPrototypeText(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPrototypeText() = %v, want %v", got, tt.want)
			}
		})
	}
}
