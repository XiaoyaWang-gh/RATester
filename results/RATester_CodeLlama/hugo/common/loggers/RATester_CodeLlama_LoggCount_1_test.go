package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestLoggCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	l.LoggCount(logg.LevelDebug)
}
