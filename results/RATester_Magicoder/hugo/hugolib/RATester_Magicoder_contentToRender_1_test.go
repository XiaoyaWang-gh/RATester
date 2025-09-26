package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/parser/pageparser"
)

func TestcontentToRender_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pi := &contentParseInfo{
		h:                       &HugoSites{},
		pid:                     1,
		sourceKey:               "test.md",
		openSource:              nil,
		frontMatter:             map[string]any{},
		hasSummaryDivider:       false,
		posMainContent:          0,
		hasNonMarkdownShortcode: false,
		itemsStep1:              pageparser.Items{},
		itemsStep2:              []any{},
	}

	source := []byte("test content")
	renderedShortcodes := map[string]shortcodeRenderer{}

	result, hasVariants, err := pi.contentToRender(ctx, source, renderedShortcodes)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result) != len(source) {
		t.Errorf("Expected result length to be %d, but got %d", len(source), len(result))
	}

	if hasVariants {
		t.Errorf("Expected hasVariants to be false, but got true")
	}
}
