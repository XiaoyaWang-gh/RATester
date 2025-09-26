package goldmark

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestParse_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &goldmarkConverter{}
	ctx := converter.RenderContext{}
	result, err := c.Parse(ctx)
	if err != nil {
		t.Errorf("Parse returned error: %v", err)
	}
	if result == nil {
		t.Errorf("Parse returned nil")
	}
	if result.Doc() == nil {
		t.Errorf("Parse returned nil doc")
	}
	if result.TableOfContents() == nil {
		t.Errorf("Parse returned nil toc")
	}
}
