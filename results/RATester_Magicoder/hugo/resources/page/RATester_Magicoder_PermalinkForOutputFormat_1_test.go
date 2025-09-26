package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/output"
)

func TestPermalinkForOutputFormat_1(t *testing.T) {
	type args struct {
		s *helpers.PathSpec
		f output.Format
	}
	tests := []struct {
		name string
		p    TargetPaths
		args args
		want string
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

			if got := tt.p.PermalinkForOutputFormat(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("PermalinkForOutputFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
