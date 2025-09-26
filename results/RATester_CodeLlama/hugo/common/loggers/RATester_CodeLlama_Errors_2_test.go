package loggers

import (
	"fmt"
	"testing"
)

func TestErrors_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	l.errors.WriteString("test")
	if l.Errors() != "test" {
		t.Errorf("Expected %s, got %s", "test", l.Errors())
	}
}
