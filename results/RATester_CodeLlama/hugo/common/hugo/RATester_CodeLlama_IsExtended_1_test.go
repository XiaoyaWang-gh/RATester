package hugo

import (
	"fmt"
	"testing"
)

func TestIsExtended_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var i HugoInfo

	// Act
	result := i.IsExtended()

	// Assert
	if result != IsExtended {
		t.Errorf("Expected %v, but got %v", IsExtended, result)
	}
}
