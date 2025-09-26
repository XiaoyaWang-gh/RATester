package goldmark

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestExtend_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &tocExtension{}
	m := goldmark.New()
	e.Extend(m)
	if m.Parser() == nil {
		t.Error("parser is nil")
	}
	if m.Renderer() == nil {
		t.Error("renderer is nil")
	}
}
