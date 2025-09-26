package utils

import (
	"fmt"
	"testing"
)

func TestSelfDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	expected := "."

	// Act
	actual := SelfDir()

	// Assert
	if actual != expected {
		t.Errorf("Expected %s, but actual is %s", expected, actual)
	}
}
