package pageparser

import (
	"fmt"
	"testing"
)

func TestIsShortcodeName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var i Item
	i.Type = tScName

	// Act
	result := i.IsShortcodeName()

	// Assert
	if !result {
		t.Errorf("Expected IsShortcodeName to return true, but got false")
	}
}
