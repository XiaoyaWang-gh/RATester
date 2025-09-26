package tableofcontents

import (
	"fmt"
	"html/template"
	"testing"
)

func TestToHTML_1(t *testing.T) {
	type testCase struct {
		name     string
		toc      *Fragments
		start    int
		stop     int
		ordered  bool
		expected template.HTML
	}

	testCases := []testCase{
		{
			name:     "nil Fragments",
			toc:      nil,
			start:    1,
			stop:     6,
			ordered:  false,
			expected: "",
		},
		{
			name:     "empty Fragments",
			toc:      &Fragments{},
			start:    1,
			stop:     6,
			ordered:  false,
			expected: "",
		},
		{
			name: "Fragments with headings",
			toc: &Fragments{
				Headings: Headings{
					&Heading{
						ID:    "heading1",
						Level: 1,
						Title: "Heading 1",
						Headings: Headings{
							&Heading{
								ID:    "subheading1",
								Level: 2,
								Title: "Subheading 1",
							},
						},
					},
				},
			},
			start:    1,
			stop:     6,
			ordered:  false,
			expected: "<ul><li><a href=\"#heading1\">Heading 1</a><ul><li><a href=\"#subheading1\">Subheading 1</a></li></ul></li></ul>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.toc.ToHTML(tc.start, tc.stop, tc.ordered)
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
