package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestNewDatadogLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	ddLogger := NewDatadogLogger(logger)

	if ddLogger == nil {
		t.Error("Expected a DatadogLogger, but got nil")
	}
}
