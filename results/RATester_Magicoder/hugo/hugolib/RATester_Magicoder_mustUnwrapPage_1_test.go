package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestmustUnwrapPage_1(t *testing.T) {
	type args struct {
		in any
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

			if got := mustUnwrapPage(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mustUnwrapPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
