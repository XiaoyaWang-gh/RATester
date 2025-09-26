package tables

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestExtend_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &ext{}
	m := goldmark.New()
	e.Extend(m)
	if m.Renderer() == nil {
		t.Error("expected non-nil renderer")
	}
}
