package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
	"text/template/parse"

	"github.com/gohugoio/hugo/tpl"
)

func TestgetParseTree_1(t *testing.T) {
	type args struct {
		templ tpl.Template
	}
	tests := []struct {
		name string
		args args
		want *parse.Tree
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

			if got := getParseTree(tt.args.templ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getParseTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
