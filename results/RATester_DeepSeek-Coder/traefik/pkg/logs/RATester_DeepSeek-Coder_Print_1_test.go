package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestPrint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	l := LogrusStdWrapper{logger: logger}

	msg := "test message"
	l.Print(msg)

	// Add your assertions here to verify the output of the Print method
}
