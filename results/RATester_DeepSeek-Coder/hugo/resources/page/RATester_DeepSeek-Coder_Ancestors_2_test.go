package page

import (
	"fmt"
	"testing"
)

func TestAncestors_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	nop := nopPage(0)

	// Act
	result := nop.Ancestors()

	// Assert
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
