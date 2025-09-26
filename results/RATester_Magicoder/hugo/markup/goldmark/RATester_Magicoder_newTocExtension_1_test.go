package goldmark

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
)

func TestnewTocExtension_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	options := []renderer.Option{
		// Add your options here
	}

	ext := newTocExtension(options)

	if ext == nil {
		t.Error("Expected newTocExtension to return a non-nil value")
	}

	if _, ok := ext.(goldmark.Extender); !ok {
		t.Error("Expected newTocExtension to return a value that implements goldmark.Extender")
	}
}
