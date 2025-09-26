package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestDebugf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	l := ElasticLogger{logger: logger}

	l.Debugf("test message")

	// assert something here
}
