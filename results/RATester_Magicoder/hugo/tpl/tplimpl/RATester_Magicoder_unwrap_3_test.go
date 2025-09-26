package tplimpl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func Testunwrap_3(t *testing.T) {
	type args struct {
		templ tpl.Template
	}
	tests := []struct {
		name string
		args args
		want tpl.Template
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

			if got := unwrap(tt.args.templ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
