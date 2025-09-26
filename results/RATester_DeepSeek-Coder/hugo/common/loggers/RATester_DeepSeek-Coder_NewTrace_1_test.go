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

	if logger.Level() != logg.LevelTrace {
		t.Errorf("Expected level to be %v, got %v", logg.LevelTrace, logger.Level())
	}

	if logger.Out() != os.Stdout {
		t.Errorf("Expected output to be %v, got %v", os.Stdout, logger.Out())
	}

	logger.Debugf("This is a %s", "debug")
	logger.Infof("This is an %s", "info")
	logger.Warnf("This is a %s", "warning")
	logger.Errorf("This is an %s", "error")

	logs := logger.Errors()
	if logs != "" {
		t.Errorf("Expected errors to be empty, got %v", logs)
	}
}
