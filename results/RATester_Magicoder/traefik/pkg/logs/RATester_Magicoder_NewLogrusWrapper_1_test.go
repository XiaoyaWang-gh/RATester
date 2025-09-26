package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestNewLogrusWrapper_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logrusWrapper := NewLogrusWrapper(logger)

	if logrusWrapper == nil {
		t.Error("Expected logrusWrapper to be not nil")
	}
}
