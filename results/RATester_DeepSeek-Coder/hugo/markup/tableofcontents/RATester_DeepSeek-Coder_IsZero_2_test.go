package tableofcontents

import (
	"fmt"
	"testing"
)

func TestIsZero_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := Heading{}
	if !h.IsZero() {
		t.Errorf("Expected IsZero to return true, got false")
	}

	h.ID = "1"
	if h.IsZero() {
		t.Errorf("Expected IsZero to return false, got true")
	}

	h.ID = ""
	h.Title = "Title"
	if h.IsZero() {
		t.Errorf("Expected IsZero to return false, got true")
	}

	h.Title = ""
	h.Headings = make(Headings, 1)
	if h.IsZero() {
		t.Errorf("Expected IsZero to return false, got true")
	}
}
