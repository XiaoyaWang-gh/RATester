package blockquotes

import (
	"fmt"
	"testing"
)

func TestnewHTMLRenderer_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := newHTMLRenderer()
	if r == nil {
		t.Error("newHTMLRenderer() should not return nil")
	}
}
