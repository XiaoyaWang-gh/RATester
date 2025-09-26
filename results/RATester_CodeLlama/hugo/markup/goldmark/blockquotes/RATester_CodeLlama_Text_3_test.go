package blockquotes

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types/hstring"
)

func TestText_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &blockquoteContext{
		text: hstring.HTML("text"),
	}
	if got := c.Text(); got != c.text {
		t.Errorf("Text() = %v, want %v", got, c.text)
	}
}
