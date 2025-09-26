package tableofcontents

import (
	"fmt"
	"strings"
	"testing"
)

func TestwriteHeading_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tocBuilder{
		s: strings.Builder{},
		h: Headings{
			{
				ID:    "test-id",
				Level: 1,
				Title: "Test Title",
				Headings: Headings{
					{
						ID:    "test-id-child",
						Level: 2,
						Title: "Test Title Child",
					},
				},
			},
		},
		startLevel: 1,
		stopLevel:  2,
		ordered:    true,
	}

	b.writeHeading(1, 0, &Heading{
		ID:    "test-id",
		Level: 1,
		Title: "Test Title",
		Headings: Headings{
			{
				ID:    "test-id-child",
				Level: 2,
				Title: "Test Title Child",
			},
		},
	})

	expected := "<li><a href=\"#test-id\">Test Title</a><li><a href=\"#test-id-child\">Test Title Child</a></li></li>\n"
	if b.s.String() != expected {
		t.Errorf("Expected %q, got %q", expected, b.s.String())
	}
}
