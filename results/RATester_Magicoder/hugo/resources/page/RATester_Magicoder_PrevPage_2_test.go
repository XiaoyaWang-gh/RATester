package page

import (
	"fmt"
	"testing"
)

func TestPrevPage_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	nopPage := nopPage(0)

	// Act
	result := nopPage.PrevPage()

	// Assert
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
