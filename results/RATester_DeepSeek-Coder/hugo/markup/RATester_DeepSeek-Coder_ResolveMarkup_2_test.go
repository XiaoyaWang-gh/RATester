package markup

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestResolveMarkup_2(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test with 'goldmark'",
			arg:  "goldmark",
			want: media.DefaultContentTypes.Markdown.SubType,
		},
		{
			name: "Test with 'asciidocext'",
			arg:  "asciidocext",
			want: media.DefaultContentTypes.AsciiDoc.SubType,
		},
		{
			name: "Test with 'unknown'",
			arg:  "unknown",
			want: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ResolveMarkup(tt.arg)
			if got != tt.want {
				t.Errorf("ResolveMarkup(%s) = %s; want %s", tt.arg, got, tt.want)
			}
		})
	}
}
