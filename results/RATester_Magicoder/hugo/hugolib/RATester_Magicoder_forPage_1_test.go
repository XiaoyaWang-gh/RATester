package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
	"github.com/gohugoio/hugo/source"
)

func TestforPage_1(t *testing.T) {
	type args struct {
		p page.Page
	}
	tests := []struct {
		name string
		g    *gitInfo
		args args
		want source.GitInfo
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

			if got := tt.g.forPage(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gitInfo.forPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
