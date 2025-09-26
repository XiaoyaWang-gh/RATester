package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/related"
)

func TestRelatedKeywords_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var (
		p = &pageMeta{
			pageMetaParams: &pageMetaParams{
				paramsOriginal: map[string]any{
					"keywords": []string{"foo", "bar", "baz"},
				},
			},
		}
		cfg = related.IndexConfig{
			Name: "keywords",
			Type: "strings",
		}
	)

	keywords, err := p.RelatedKeywords(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if len(keywords) != 3 {
		t.Errorf("Expected 3 keywords, got %d", len(keywords))
	}

	if keywords[0].String() != "foo" {
		t.Errorf("Expected keyword %q, got %q", "foo", keywords[0].String())
	}

	if keywords[1].String() != "bar" {
		t.Errorf("Expected keyword %q, got %q", "bar", keywords[1].String())
	}

	if keywords[2].String() != "baz" {
		t.Errorf("Expected keyword %q, got %q", "baz", keywords[2].String())
	}
}
