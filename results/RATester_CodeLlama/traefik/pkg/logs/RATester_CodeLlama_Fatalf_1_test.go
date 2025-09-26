package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestFatalf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := LogrusStdWrapper{
		logger: zerolog.New(os.Stdout),
	}
	l.Fatalf("test %s", "test")
}
