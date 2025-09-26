package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestNewRetryableHTTPLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	retryableLogger := NewRetryableHTTPLogger(logger)

	if retryableLogger == nil {
		t.Error("Expected a RetryableHTTPLogger, but got nil")
	}
}
