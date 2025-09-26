package loggers

import (
	"fmt"
	"testing"
)

func TestnewLogLevelCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	counter := newLogLevelCounter()
	if counter == nil {
		t.Error("Expected a non-nil logLevelCounter, got nil")
	}
	if counter.counters == nil {
		t.Error("Expected a non-nil counters map, got nil")
	}
}
