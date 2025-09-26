package page

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestresolveParagraphTagAndSetWrapper_1(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		mt       media.Type
		expected tagReStartEnd
	}{
		{
			name:   "Test case 1",
			source: "<div class=\"document\">Hello, World!</div>",
			mt:     media.Type{SubType: media.DefaultContentTypes.ReStructuredText.SubType},
			expected: tagReStartEnd{
				startEndOfString: regexp.MustCompile(`<div class="document">`),
				endEndOfString:   regexp.MustCompile(`</div>`),
				tagName:          "div",
			},
		},
		{
			name:   "Test case 2",
			source: "Hello, World!",
			mt:     media.Type{SubType: "other"},
			expected: tagReStartEnd{
				startEndOfString: regexp.MustCompile(`<p>`),
				endEndOfString:   regexp.MustCompile(`</p>`),
				tagName:          "p",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			summary := HtmlSummary{
				source: test.source,
			}
			result := summary.resolveParagraphTagAndSetWrapper(test.mt)

			if result.startEndOfString.String() != test.expected.startEndOfString.String() ||
				result.endEndOfString.String() != test.expected.endEndOfString.String() ||
				result.tagName != test.expected.tagName {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
