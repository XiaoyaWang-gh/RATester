package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	l := OxyWrapper{logger: logger}

	l.Error("test error")
}
