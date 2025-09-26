package loggers

import (
	"fmt"
	"testing"
)

func TestNewDefault_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	logger := NewDefault()
	// Assert
	if logger == nil {
		t.Errorf("NewDefault() returned nil")
	}
}
