package tableofcontents

import (
	"fmt"
	"strings"
	"testing"
)

func TestwriteHeadings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &tocBuilder{
		s: strings.Builder{},
		h: Headings{
			{
				Level: 1,
				Headings: Headings{
					{
						Level: 2,
						Headings: Headings{
							{
								Level: 3,
							},
						},
					},
				},
			},
		},
		startLevel: 1,
		stopLevel:  -1,
		ordered:    false,
	}

	b.writeHeadings(1, 0, b.h)

	expected := "<ul>\n<li>\n<ul>\n<li>\n<ul>\n<li>\n</li>\n</ul>\n</li>\n</ul>\n</li>\n</ul>\n"
	if b.s.String() != expected {
		t.Errorf("Expected %q, got %q", expected, b.s.String())
	}
}
