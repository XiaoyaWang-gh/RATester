package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestPrintln_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var args []interface{}
	l := LogrusStdWrapper{
		logger: zerolog.New(os.Stdout),
	}

	// when
	l.Println(args...)

	// then
	// TODO
}
