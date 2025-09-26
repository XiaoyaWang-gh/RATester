package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestNewWasmLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel)
	wasmLogger := NewWasmLogger(&logger)

	if wasmLogger.logger != &logger {
		t.Errorf("Expected logger to be %v, but got %v", &logger, wasmLogger.logger)
	}
}
