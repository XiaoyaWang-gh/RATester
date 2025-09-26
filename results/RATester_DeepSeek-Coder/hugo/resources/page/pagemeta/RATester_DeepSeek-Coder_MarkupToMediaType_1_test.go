package pagemeta

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestMarkupToMediaType_1(t *testing.T) {
	tests := []struct {
		name       string
		markup     string
		mediaTypes media.Types
		want       media.Type
	}{
		{
			name:   "test1",
			markup: "test1",
			mediaTypes: media.Types{
				{
					MainType:    "text",
					SubType:     "html",
					SuffixesCSV: "html",
				},
				{
					MainType:    "text",
					SubType:     "markdown",
					SuffixesCSV: "md",
				},
			},
			want: media.Type{
				MainType:    "text",
				SubType:     "html",
				SuffixesCSV: "html",
			},
		},
		{
			name:   "test2",
			markup: "test2.md",
			mediaTypes: media.Types{
				{
					MainType:    "text",
					SubType:     "html",
					SuffixesCSV: "html",
				},
				{
					MainType:    "text",
					SubType:     "markdown",
					SuffixesCSV: "md",
				},
			},
			want: media.Type{
				MainType:    "text",
				SubType:     "markdown",
				SuffixesCSV: "md",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MarkupToMediaType(tt.markup, tt.mediaTypes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarkupToMediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}
