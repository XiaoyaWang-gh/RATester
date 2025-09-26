package loggers

import (
	"fmt"
	"testing"
)

func TestNewLogLevelCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := newLogLevelCounter()
	// Assert
	if result == nil {
		t.Errorf("newLogLevelCounter() failed")
	}
}
