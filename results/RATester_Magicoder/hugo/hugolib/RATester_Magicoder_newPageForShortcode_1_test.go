package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestnewPageForShortcode_1(t *testing.T) {
	type args struct {
		p *pageState
	}
	tests := []struct {
		name string
		args args
		want page.Page
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

			if got := newPageForShortcode(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPageForShortcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
