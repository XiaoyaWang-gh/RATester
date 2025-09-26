package loggers

import (
	"fmt"
	"testing"
)

func TestLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	if got := l.Logger(); got != l.logger {
		t.Errorf("Logger() = %v, want %v", got, l.logger)
	}
}
