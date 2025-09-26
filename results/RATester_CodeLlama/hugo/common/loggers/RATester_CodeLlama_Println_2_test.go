package loggers

import (
	"fmt"
	"testing"
)

func TestPrintln_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	l.Println("test")
	if l.errors.String() != "test\n" {
		t.Errorf("expected %q, got %q", "test\n", l.errors.String())
	}
}
