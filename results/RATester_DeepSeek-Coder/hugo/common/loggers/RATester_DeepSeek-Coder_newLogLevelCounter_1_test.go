package loggers

import (
	"fmt"
	"testing"
)

func TestNewLogLevelCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	counter := newLogLevelCounter()

	if counter == nil {
		t.Errorf("Expected a non-nil logLevelCounter, got nil")
	}

	if counter.counters == nil {
		t.Errorf("Expected a non-nil counters map, got nil")
	}
}
