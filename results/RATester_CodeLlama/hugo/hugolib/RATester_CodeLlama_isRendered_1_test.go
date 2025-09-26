package hugolib

import (
	"fmt"
	"testing"
)

func TestIsRendered_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	po := &pageOutput{}
	// Act
	result := po.isRendered()
	// Assert
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
