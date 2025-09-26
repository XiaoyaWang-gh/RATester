package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestWarn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	l := OxyWrapper{logger: logger}

	testString := "test"
	testArgs := []interface{}{"test"}

	l.Warn(testString, testArgs...)

	// Add assertions here to check if the expected log was written
}
