package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/output"
	"github.com/gohugoio/hugo/resources/kinds"
)

func TestCreateDefaultOutputFormats_1(t *testing.T) {
	tests := []struct {
		name       string
		allFormats output.Formats
		want       map[string][]string
	}{
		{
			name:       "no formats",
			allFormats: output.Formats{},
			want:       nil,
		},
		{
			name: "only html",
			allFormats: output.Formats{
				{Name: "HTML"},
			},
			want: map[string][]string{
				kinds.KindPage:     {"HTML"},
				kinds.KindHome:     {"HTML"},
				kinds.KindSection:  {"HTML"},
				kinds.KindTerm:     {"HTML"},
				kinds.KindTaxonomy: {"HTML"},
			},
		},
		{
			name: "html and rss",
			allFormats: output.Formats{
				{Name: "HTML"},
				{Name: "RSS"},
			},
			want: map[string][]string{
				kinds.KindPage:     {"HTML"},
				kinds.KindHome:     {"HTML", "RSS"},
				kinds.KindSection:  {"HTML", "RSS"},
				kinds.KindTerm:     {"HTML"},
				kinds.KindTaxonomy: {"HTML"},
				"rss":              {"RSS"},
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
