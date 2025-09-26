package pageparser

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
)

func TestFormatFromFrontMatterType_1(t *testing.T) {
	tests := []struct {
		name string
		typ  ItemType
		want metadecoders.Format
	}{
		{
			name: "JSON",
			typ:  TypeFrontMatterJSON,
			want: metadecoders.JSON,
		},
		{
			name: "ORG",
			typ:  TypeFrontMatterORG,
			want: metadecoders.ORG,
		},
		{
			name: "TOML",
			typ:  TypeFrontMatterTOML,
			want: metadecoders.TOML,
		},
		{
			name: "YAML",
			typ:  TypeFrontMatterYAML,
			want: metadecoders.YAML,
		},
		{
			name: "default",
			typ:  ItemType(0),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FormatFromFrontMatterType(tt.typ); got != tt.want {
				t.Errorf("FormatFromFrontMatterType() = %v, want %v", got, tt.want)
			}
		})
	}
}
