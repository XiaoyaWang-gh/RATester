package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestNewTrace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	logger := NewTrace()
	// Assert
	if logger.Level() != logg.LevelTrace {
		t.Errorf("Expected %d, got %d", logg.LevelTrace, logger.Level())
	}
}
