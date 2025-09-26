package tplimpl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func TestfindTemplateIn_1(t *testing.T) {
	type args struct {
		name string
		in   tpl.Template
	}
	tests := []struct {
		name  string
		args  args
		want  tpl.Template
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := findTemplateIn(tt.args.name, tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTemplateIn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findTemplateIn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
