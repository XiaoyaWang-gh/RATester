package loggers

import (
	"fmt"
	"os"
	"testing"

	"github.com/bep/logg"
)

func TestNewTrace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := NewTrace()

	if logger == nil {
		t.Error("Expected logger to be not nil")
	}

	if logger.Level() != logg.LevelTrace {
		t.Errorf("Expected logger level to be %v, but got %v", logg.LevelTrace, logger.Level())
	}

	if logger.Out() != os.Stdout {
		t.Errorf("Expected logger output to be %v, but got %v", os.Stdout, logger.Out())
	}

	// Add more test cases as needed
}
