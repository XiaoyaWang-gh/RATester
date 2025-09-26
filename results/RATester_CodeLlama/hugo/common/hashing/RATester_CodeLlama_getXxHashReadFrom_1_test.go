package hashing

import (
	"fmt"
	"testing"
)

func TestGetXxHashReadFrom_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := getXxHashReadFrom()
	// Assert
	if result == nil {
		t.Errorf("getXxHashReadFrom() failed")
	}
}
