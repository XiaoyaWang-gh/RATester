package layouts

import (
	"fmt"
	"testing"
)

func TestAddLayoutVariations_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &layoutBuilder{}
	l.addLayoutVariations("baseof")
	if len(l.layoutVariations) != 1 {
		t.Errorf("Expected 1 layout variation, got %d", len(l.layoutVariations))
	}
	if l.layoutVariations[0] != "baseof-baseof" {
		t.Errorf("Expected layout variation to be 'baseof-baseof', got %s", l.layoutVariations[0])
	}
}
