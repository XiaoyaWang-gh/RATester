package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestErrorf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	l := ElasticLogger{logger: logger}

	format := "test error"
	args := []interface{}{"arg1", "arg2"}

	l.Errorf(format, args...)

	// Add assertions here to verify the behavior of the method under test
}
