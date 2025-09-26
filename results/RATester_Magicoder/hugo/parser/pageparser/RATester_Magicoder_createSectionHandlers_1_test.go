package pageparser

import (
	"fmt"
	"testing"
)

func TestcreateSectionHandlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("[shortcode]"),
		cfg: Config{
			NoFrontMatter:    false,
			NoSummaryDivider: false,
		},
	}

	handlers := createSectionHandlers(l)

	if handlers == nil {
		t.Errorf("Expected handlers to be created, but got nil")
	}

	if len(handlers.handlers) != 2 {
		t.Errorf("Expected 2 handlers, but got %d", len(handlers.handlers))
	}

	if handlers.handlers[0].l != l {
		t.Errorf("Expected handler's lexer to be the same as the input lexer")
	}

	if handlers.handlers[1].l != l {
		t.Errorf("Expected handler's lexer to be the same as the input lexer")
	}
}
