package loggers

import (
	"fmt"
	"testing"
)

func TestWarn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	l.Warn()
}
