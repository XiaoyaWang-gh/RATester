package loggers

import (
	"fmt"
	"testing"
)

func TestErrorln_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	l.Errorln("test")
	if l.errors.String() != "test\n" {
		t.Errorf("Expected %q, got %q", "test\n", l.errors.String())
	}
}
