package gin

import (
	"fmt"
	"testing"
)

func TestIsDebugging_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	ginMode = debugCode

	// Act
	result := IsDebugging()

	// Assert
	if !result {
		t.Errorf("Expected IsDebugging() to return true, but got false")
	}
}
