package page

import (
	"fmt"
	"testing"
)

func TestByWeight_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var p Pages

	// Act
	var pages Pages
	pages = p.ByWeight()

	// Assert
	if len(pages) != 0 {
		t.Errorf("Expected pages to be empty, but got %d pages", len(pages))
	}
}
