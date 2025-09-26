package hugolib

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/tableofcontents"
)

func TestHeadingsFiltered_1(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name     string
		headings tableofcontents.Headings
		expected tableofcontents.Headings
	}{
		{
			name: "no headings",
			headings: tableofcontents.Headings{
				{Level: 1, Title: "Heading 1"},
				{Level: 2, Title: "Heading 2"},
			},
			expected: tableofcontents.Headings{
				{Level: 1, Title: "Heading 1"},
				{Level: 2, Title: "Heading 2"},
			},
		},
		{
			name: "filter headings",
			headings: tableofcontents.Headings{
				{Level: 1, Title: "Heading 1"},
				{Level: 2, Title: "Heading 2"},
			},
			expected: tableofcontents.Headings{
				{Level: 1, Title: "Heading 1"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &pageHeadingsFiltered{
				headings: tc.headings,
			}

			result := p.HeadingsFiltered(ctx)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
