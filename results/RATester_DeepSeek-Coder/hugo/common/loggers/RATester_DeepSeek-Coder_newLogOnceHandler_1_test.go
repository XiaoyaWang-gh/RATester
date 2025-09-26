package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestNewLogOnceHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	threshold := logg.LevelError
	h := newLogOnceHandler(threshold)

	if h.threshold != threshold {
		t.Errorf("Expected threshold to be %v, got %v", threshold, h.threshold)
	}

	if h.seen == nil {
		t.Errorf("Expected seen to be initialized")
	}
}
