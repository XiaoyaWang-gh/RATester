package loggers

import (
	"fmt"
	"testing"
)

func TestError_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	l.Error()
}
