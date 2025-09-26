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

	h := Headings{
		{
			ID:    "h1",
			Level: 1,
			Title: "h1",
			Headings: Headings{
				{
					ID:    "h1-1",
					Level: 2,
					Title: "h1-1",
				},
				{
					ID:    "h1-2",
					Level: 2,
					Title: "h1-2",
				},
			},
		},
		{
			ID:    "h2",
			Level: 1,
			Title: "h2",
			Headings: Headings{
				{
					ID:    "h2-1",
					Level: 2,
					Title: "h2-1",
				},
				{
					ID:    "h2-2",
					Level: 2,
					Title: "h2-2",
				},
			},
		},
	}

	fn := func(h *Heading) bool {
		return h.Level == 2
	}

	out := h.FilterBy(fn)

	if len(out) != 2 {
		t.Errorf("expected 2 headings, got %d", len(out))
	}

	if out[0].ID != "h1-1" {
		t.Errorf("expected h1-1, got %s", out[0].ID)
	}

	if out[1].ID != "h2-1" {
		t.Errorf("expected h2-1, got %s", out[1].ID)
	}
}
