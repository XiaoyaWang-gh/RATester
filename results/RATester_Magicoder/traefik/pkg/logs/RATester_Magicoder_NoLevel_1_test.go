package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestNoLevel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	level := zerolog.DebugLevel
	newLogger := NoLevel(logger, level)

	if newLogger.GetLevel() != level {
		t.Errorf("Expected level to be %v, but got %v", level, newLogger.GetLevel())
	}
}
