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
		t.Error("Expected IsZero to return true for an empty Heading")
	}

	h.ID = "1"
	if h.IsZero() {
		t.Error("Expected IsZero to return false for a Heading with ID")
	}

	h.ID = ""
	h.Title = "Title"
	if h.IsZero() {
		t.Error("Expected IsZero to return false for a Heading with Title")
	}

	h.Title = ""
	if !h.IsZero() {
		t.Error("Expected IsZero to return true for a Heading with no ID or Title")
	}
}
