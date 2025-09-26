package commands

import (
	"fmt"
	"testing"
)

func TestNewConfigCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := newConfigCommand()
	// Assert
	if result == nil {
		t.Errorf("newConfigCommand() = nil, want not nil")
	}
}
