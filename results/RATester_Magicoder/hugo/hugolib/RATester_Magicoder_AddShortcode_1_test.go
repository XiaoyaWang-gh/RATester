package hugolib

import (
	"fmt"
	"testing"
)

func TestAddShortcode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := &contentParseInfo{
		itemsStep2: make([]any, 0),
	}

	shortcode := &shortcode{
		name: "test",
		inner: []any{
			"test inner",
		},
	}

	info.AddShortcode(shortcode)

	if len(info.itemsStep2) != 1 {
		t.Errorf("Expected itemsStep2 to have 1 item, but got %d", len(info.itemsStep2))
	}

	if info.hasNonMarkdownShortcode != shortcode.insertPlaceholder() {
		t.Errorf("Expected hasNonMarkdownShortcode to be %v, but got %v", shortcode.insertPlaceholder(), info.hasNonMarkdownShortcode)
	}
}
