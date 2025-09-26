package highlight

import (
	"fmt"
	"html/template"
	"testing"
)

func TestInner_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := HighlightResult{
		innerLow:    1,
		innerHigh:   2,
		highlighted: template.HTML("foo"),
	}
	if got := h.Inner(); got != template.HTML("oo") {
		t.Errorf("Inner() = %v, want %v", got, template.HTML("oo"))
	}
}
