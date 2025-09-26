package tableofcontents

import (
	"fmt"
	"testing"
)

func TestWalk_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	toc := Fragments{
		Headings: Headings{
			{
				ID:    "h1",
				Level: 1,
				Title: "h1",
				Headings: Headings{
					{
						ID:    "h2",
						Level: 2,
						Title: "h2",
						Headings: Headings{
							{
								ID:    "h3",
								Level: 3,
								Title: "h3",
							},
						},
					},
				},
			},
		},
	}
	fn := func(h *Heading) {
		if h.ID != "h3" {
			t.Errorf("expected h3, got %s", h.ID)
		}
	}
	toc.walk(fn)
}
