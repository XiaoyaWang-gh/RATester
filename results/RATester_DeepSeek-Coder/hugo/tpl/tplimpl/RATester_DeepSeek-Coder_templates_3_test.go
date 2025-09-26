package tplimpl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func TestTemplates_3(t *testing.T) {
	type args struct {
		in tpl.Template
	}
	tests := []struct {
		name string
		args args
		want []tpl.Template
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

			if got := templates(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templates() = %v, want %v", got, tt.want)
			}
		})
	}
}
