package highlight

import (
	"fmt"
	"html/template"
	"testing"
)

func TestWrapped_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := HighlightResult{
		highlighted: template.HTML("foo"),
	}
	if got := h.Wrapped(); got != h.highlighted {
		t.Errorf("Wrapped() = %v, want %v", got, h.highlighted)
	}
}
