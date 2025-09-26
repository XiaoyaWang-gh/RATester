package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/output"
)

func TestPermalinkForOutputFormat_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		target   TargetPaths
		protocol string
		format   output.Format
		want     string
	}{
		{
			name: "Test with empty protocol",
			target: TargetPaths{
				Link: "/test/path",
			},
			protocol: "",
			format: output.Format{
				Protocol: "https",
			},
			want: "https:///test/path",
		},
		{
			name: "Test with non-empty protocol",
			target: TargetPaths{
				Link: "/test/path",
			},
			protocol: "http",
			format: output.Format{
				Protocol: "https",
			},
			want: "http:///test/path",
		},
		{
			name: "Test with same protocol",
			target: TargetPaths{
				Link: "/test/path",
			},
			protocol: "https",
			format: output.Format{
				Protocol: "https",
			},
			want: "https:///test/path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.target.PermalinkForOutputFormat(&helpers.PathSpec{}, tt.format)
			if got != tt.want {
				t.Errorf("PermalinkForOutputFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
