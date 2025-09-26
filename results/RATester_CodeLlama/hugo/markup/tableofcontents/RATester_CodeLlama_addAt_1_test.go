package tableofcontents

import (
	"fmt"
	"testing"
)

func TestAddAt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	toc := &Fragments{}
	h := &Heading{
		ID:    "id",
		Level: 1,
		Title: "title",
	}
	row := 0
	level := 0

	toc.addAt(h, row, level)

	if len(toc.Headings) != 1 {
		t.Errorf("Expected toc.Headings to have length 1, but was %d", len(toc.Headings))
	}

	if toc.Headings[0].ID != "id" {
		t.Errorf("Expected toc.Headings[0].ID to be %q, but was %q", "id", toc.Headings[0].ID)
	}

	if toc.Headings[0].Level != 1 {
		t.Errorf("Expected toc.Headings[0].Level to be %d, but was %d", 1, toc.Headings[0].Level)
	}

	if toc.Headings[0].Title != "title" {
		t.Errorf("Expected toc.Headings[0].Title to be %q, but was %q", "title", toc.Headings[0].Title)
	}
}
