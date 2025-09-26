package loggers

import (
	"fmt"
	"testing"
)

func TestSprint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	v := []any{"a", "b", "c"}
	l.sprint(v...)
}
