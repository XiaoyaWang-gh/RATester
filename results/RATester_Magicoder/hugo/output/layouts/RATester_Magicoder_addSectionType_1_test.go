package layouts

import (
	"fmt"
	"testing"
)

func TestaddSectionType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &layoutBuilder{
		d: LayoutDescriptor{
			Section: "testSection",
		},
	}
	l.addSectionType()
	if len(l.typeVariations) != 1 {
		t.Errorf("Expected 1 typeVariation, got %d", len(l.typeVariations))
	}
	if l.typeVariations[0] != "testSection" {
		t.Errorf("Expected typeVariation to be 'testSection', got '%s'", l.typeVariations[0])
	}
}
