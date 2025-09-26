package page

import (
	"fmt"
	"testing"
)

func TestByDate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var p Pages

	// Act
	var result Pages
	result = p.ByDate()

	// Assert
	if len(result) != 0 {
		t.Errorf("Expected result to be empty, but got %v", result)
	}
}
