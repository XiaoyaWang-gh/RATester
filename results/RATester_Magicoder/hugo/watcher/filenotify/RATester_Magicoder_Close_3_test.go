package filenotify

import (
	"fmt"
	"testing"
)

func TestClose_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &filePoller{
		watches: make(map[string]struct{}),
		done:    make(chan struct{}),
	}

	err := w.Close()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !w.closed {
		t.Error("Expected closed to be true")
	}

	if w.done == nil {
		t.Error("Expected done to be non-nil")
	}

	if len(w.watches) != 0 {
		t.Errorf("Expected watches to be empty, got %v", w.watches)
	}
}
