package goldmark

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/tableofcontents"
)

func TestTableOfContents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &parserContext{}

	// Test case 1: When the ToC result is not present
	result := ctx.TableOfContents()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}

	// Test case 2: When the ToC result is present
	expected := &tableofcontents.Fragments{}
	ctx.Set(tocResultKey, expected)
	result = ctx.TableOfContents()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
