package hugolib

import (
	"fmt"
	"testing"
)

func TestShouldListAny_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	p := &pageMeta{}
	// Act
	result := p.shouldListAny()
	// Assert
	if result != true {
		t.Errorf("shouldListAny() = %v, want %v", result, true)
	}
}
