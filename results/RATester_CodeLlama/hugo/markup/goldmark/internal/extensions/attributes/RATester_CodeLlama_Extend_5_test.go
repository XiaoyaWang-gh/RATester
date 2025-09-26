package attributes

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestExtend_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &attrExtension{}
	m := goldmark.New()
	a.Extend(m)
	if m.Parser() == nil {
		t.Error("parser is nil")
	}
	if m.Renderer() == nil {
		t.Error("renderer is nil")
	}
}
