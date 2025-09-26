package highlight

import (
	"fmt"
	"testing"
)

func TestIsDefaultCodeBlockRenderer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := chromaHighlighter{}

	if !h.IsDefaultCodeBlockRenderer() {
		t.Errorf("Expected IsDefaultCodeBlockRenderer to return true")
	}
}
