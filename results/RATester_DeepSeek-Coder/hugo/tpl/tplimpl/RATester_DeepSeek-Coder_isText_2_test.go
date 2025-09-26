package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl"
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
		// TODO: Add test cases.
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
