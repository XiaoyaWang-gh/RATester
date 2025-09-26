package tplimpl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func TestnewStandaloneTextTemplate_1(t *testing.T) {
	type args struct {
		funcs map[string]any
	}
	tests := []struct {
		name string
		args args
		want tpl.TemplateParseFinder
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

			if got := newStandaloneTextTemplate(tt.args.funcs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStandaloneTextTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
