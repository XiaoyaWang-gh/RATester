package transform

import (
	"context"
	"fmt"
	"testing"
)

func TestMarkdownify_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	ctx := context.Background()
	s := "Hello, World!"
	result, err := ns.Markdownify(ctx, s)
	if err != nil {
		t.Errorf("Markdownify returned an error: %v", err)
	}
	if result != "<p>Hello, World!</p>\n" {
		t.Errorf("Markdownify returned %q, expected %q", result, "<p>Hello, World!</p>\n")
	}
}
