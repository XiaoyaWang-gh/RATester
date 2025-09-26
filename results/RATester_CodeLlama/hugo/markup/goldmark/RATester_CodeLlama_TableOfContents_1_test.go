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

	p := &parserContext{}
	p.Set(tocResultKey, &tableofcontents.Fragments{})
	if p.TableOfContents() == nil {
		t.Errorf("TableOfContents() = nil, want not nil")
	}
}
