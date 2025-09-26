package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	types := Types{
		{Type: "application/rss+xml"},
		{Type: "text/html"},
		{Type: "text/markdown"},
		{Type: "text/asciidoc"},
		{Type: "text/pandoc"},
		{Type: "text/restructuredtext"},
		{Type: "text/orgmode"},
	}

	contentTypes := ContentTypes{
		HTML:             Type{Type: "text/html"},
		Markdown:         Type{Type: "text/markdown"},
		AsciiDoc:         Type{Type: "text/asciidoc"},
		Pandoc:           Type{Type: "text/pandoc"},
		ReStructuredText: Type{Type: "text/restructuredtext"},
		EmacsOrgMode:     Type{Type: "text/orgmode"},
	}

	expected := ContentTypes{
		HTML:             types[1],
		Markdown:         types[2],
		AsciiDoc:         types[3],
		Pandoc:           types[4],
		ReStructuredText: types[5],
		EmacsOrgMode:     types[6],
	}

	result := contentTypes.FromTypes(types)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
