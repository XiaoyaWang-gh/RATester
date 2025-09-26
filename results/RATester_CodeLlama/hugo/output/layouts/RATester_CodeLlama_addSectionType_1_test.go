package layouts

import (
	"fmt"
	"testing"
)

func TestAddSectionType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &layoutBuilder{}
	l.addSectionType()
	if len(l.typeVariations) != 0 {
		t.Errorf("Expected 0 type variations, got %d", len(l.typeVariations))
	}
	l.d.Section = "section"
	l.addSectionType()
	if len(l.typeVariations) != 1 {
		t.Errorf("Expected 1 type variations, got %d", len(l.typeVariations))
	}
	if l.typeVariations[0] != "section" {
		t.Errorf("Expected type variation to be 'section', got '%s'", l.typeVariations[0])
	}
}
