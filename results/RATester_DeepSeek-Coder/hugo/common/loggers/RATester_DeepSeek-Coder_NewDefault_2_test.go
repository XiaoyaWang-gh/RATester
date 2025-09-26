package loggers

import (
	"fmt"
	"os"
	"testing"

	"github.com/bep/logg"
)

func TestNewDefault_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := NewDefault()

	if logger == nil {
		t.Errorf("Expected a non-nil logger, got nil")
	}

	if logger.Level() != logg.LevelWarn {
		t.Errorf("Expected level to be %v, got %v", logg.LevelWarn, logger.Level())
	}

	if logger.Out() != os.Stdout {
		t.Errorf("Expected output to be %v, got %v", os.Stdout, logger.Out())
	}
}
