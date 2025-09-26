package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/output"
	"github.com/gohugoio/hugo/resources/kinds"
)

func TestcreateDefaultOutputFormats_1(t *testing.T) {
	tests := []struct {
		name       string
		allFormats output.Formats
		want       map[string][]string
	}{
		{
			name:       "no output formats",
			allFormats: output.Formats{},
			want:       map[string][]string{},
		},
		{
			name: "one output format",
			allFormats: output.Formats{
				{Name: "html"},
			},
			want: map[string][]string{
				kinds.KindPage:     {"html"},
				kinds.KindHome:     {"html"},
				kinds.KindSection:  {"html"},
				kinds.KindTerm:     {"html"},
				kinds.KindTaxonomy: {"html"},
			},
		},
		{
			name: "two output formats",
			allFormats: output.Formats{
				{Name: "html"},
				{Name: "rss"},
			},
			want: map[string][]string{
				kinds.KindPage:     {"html"},
				kinds.KindHome:     {"html", "rss"},
				kinds.KindSection:  {"html", "rss"},
				kinds.KindTerm:     {"html", "rss"},
				kinds.KindTaxonomy: {"html", "rss"},
				"rss":              {"rss"},
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

			got := createDefaultOutputFormats(tt.allFormats)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createDefaultOutputFormats() = %v, want %v", got, tt.want)
			}
		})
	}
}
