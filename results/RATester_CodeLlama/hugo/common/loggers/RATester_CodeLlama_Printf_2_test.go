package loggers

import (
	"fmt"
	"testing"
)

func TestPrintf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	l.Printf("test")
	if l.errors.String() != "test\n" {
		t.Errorf("l.errors.String() = %q, want %q", l.errors.String(), "test\n")
	}
}
