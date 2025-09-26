package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestnewLogOnceHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	threshold := logg.Level(1)
	handler := newLogOnceHandler(threshold)

	if handler.threshold != threshold {
		t.Errorf("Expected threshold to be %v, but got %v", threshold, handler.threshold)
	}

	if len(handler.seen) != 0 {
		t.Errorf("Expected seen map to be empty, but got %v", handler.seen)
	}
}
