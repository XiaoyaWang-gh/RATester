package media

import (
	"fmt"
	"testing"
)

func TestTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ct := ContentTypes{
		HTML:             Type{MainType: "text", SubType: "html"},
		Markdown:         Type{MainType: "text", SubType: "markdown"},
		AsciiDoc:         Type{MainType: "text", SubType: "asciidoc"},
		Pandoc:           Type{MainType: "text", SubType: "pandoc"},
		ReStructuredText: Type{MainType: "text", SubType: "rest"},
		EmacsOrgMode:     Type{MainType: "text", SubType: "org"},
	}

	ct.init()

	if ct.HTML.Type != "text/html" {
		t.Errorf("Expected HTML type to be 'text/html', got %s", ct.HTML.Type)
	}

	if ct.Markdown.Type != "text/markdown" {
		t.Errorf("Expected Markdown type to be 'text/markdown', got %s", ct.Markdown.Type)
	}

	if ct.AsciiDoc.Type != "text/asciidoc" {
		t.Errorf("Expected AsciiDoc type to be 'text/asciidoc', got %s", ct.AsciiDoc.Type)
	}

	if ct.Pandoc.Type != "text/pandoc" {
		t.Errorf("Expected Pandoc type to be 'text/pandoc', got %s", ct.Pandoc.Type)
	}

	if ct.ReStructuredText.Type != "text/rest" {
		t.Errorf("Expected ReStructuredText type to be 'text/rest', got %s", ct.ReStructuredText.Type)
	}

	if ct.EmacsOrgMode.Type != "text/org" {
		t.Errorf("Expected EmacsOrgMode type to be 'text/org', got %s", ct.EmacsOrgMode.Type)
	}
}
