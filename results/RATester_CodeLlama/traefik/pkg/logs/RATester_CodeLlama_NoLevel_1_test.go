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

	logger := zerolog.New(os.Stdout)
	logger = logger.Level(zerolog.DebugLevel)
	logger = NoLevel(logger, zerolog.InfoLevel)
	if logger.GetLevel() != zerolog.DebugLevel {
		t.Errorf("logger.GetLevel() = %v, want %v", logger.GetLevel(), zerolog.DebugLevel)
	}
}
