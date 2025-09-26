package tableofcontents

import (
	"fmt"
	"testing"
)

func TestFilterBy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h1 := &Heading{ID: "h1", Level: 1, Title: "Heading 1"}
	h2 := &Heading{ID: "h2", Level: 2, Title: "Heading 2"}
	h3 := &Heading{ID: "h3", Level: 3, Title: "Heading 3"}
	h4 := &Heading{ID: "h4", Level: 4, Title: "Heading 4"}
	h5 := &Heading{ID: "h5", Level: 5, Title: "Heading 5"}
	h6 := &Heading{ID: "h6", Level: 6, Title: "Heading 6"}

	headings := Headings{h1, h2, h3, h4, h5, h6}

	filtered := headings.FilterBy(func(h *Heading) bool {
		return h.Level <= 3
	})

	if len(filtered) != 3 {
		t.Errorf("Expected 3 headings, got %d", len(filtered))
	}

	for _, h := range filtered {
		if h.Level > 3 {
			t.Errorf("Expected headings with level <= 3, got level %d", h.Level)
		}
	}
}
