package blockquotes

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types/hstring"
)

func TestAlertTitle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &blockquoteContext{
		alert: blockQuoteAlert{
			title: "Test Title",
		},
	}

	expected := hstring.HTML("Test Title")
	result := ctx.AlertTitle()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
