package layouts

import (
	"fmt"
	"testing"
)

func TestAddTypeVariations_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &layoutBuilder{}
	l.addTypeVariations("a", "b", "c")
	if len(l.typeVariations) != 3 {
		t.Errorf("Expected 3 type variations, got %d", len(l.typeVariations))
	}
	if l.typeVariations[0] != "a" {
		t.Errorf("Expected type variation a, got %s", l.typeVariations[0])
	}
	if l.typeVariations[1] != "b" {
		t.Errorf("Expected type variation b, got %s", l.typeVariations[1])
	}
	if l.typeVariations[2] != "c" {
		t.Errorf("Expected type variation c, got %s", l.typeVariations[2])
	}
}
