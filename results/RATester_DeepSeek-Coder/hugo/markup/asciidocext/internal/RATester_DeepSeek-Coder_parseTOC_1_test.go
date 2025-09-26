package internal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/tableofcontents"
	"golang.org/x/net/html"
)

func TestParseTOC_1(t *testing.T) {
	type args struct {
		doc *html.Node
	}
	tests := []struct {
		name string
		args args
		want *tableofcontents.Fragments
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

			if got := parseTOC(tt.args.doc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTOC() = %v, want %v", got, tt.want)
			}
		})
	}
}
