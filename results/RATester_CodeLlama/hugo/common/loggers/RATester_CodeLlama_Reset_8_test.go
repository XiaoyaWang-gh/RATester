package loggers

import (
	"fmt"
	"testing"
)

func TestReset_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	l.reset = func() {}
	l.Reset()
}
