package middleware

import (
	"fmt"
	"testing"
)

func TestGenerator_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	expected := "12345678901234567890123456789012"

	// Act
	actual := generator()

	// Assert
	if actual != expected {
		t.Errorf("Expected %s, but actual is %s", expected, actual)
	}
}
