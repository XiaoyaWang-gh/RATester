package hugolib

import (
	"fmt"
	"testing"
)

func TestSectionsPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	p := pageTree{}
	// Act
	actual := p.SectionsPath()
	// Assert
	expected := ""
	if actual != expected {
		t.Errorf("SectionsPath() = %v, want %v", actual, expected)
	}
}
