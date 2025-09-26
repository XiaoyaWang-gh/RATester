package codeblocks

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestExtend_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &codeBlocksExtension{}
	m := goldmark.New()
	e.Extend(m)
	if m.Renderer() == nil {
		t.Error("expected non-nil renderer")
	}
}
