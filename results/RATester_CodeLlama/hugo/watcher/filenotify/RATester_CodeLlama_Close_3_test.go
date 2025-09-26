package filenotify

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

func TestClose_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &filePoller{
		interval: time.Second,
		watches:  make(map[string]struct{}),
		done:     make(chan struct{}),
		events:   make(chan fsnotify.Event),
		errors:   make(chan error),
	}

	w.Add("test")
	w.Close()

	if len(w.watches) != 0 {
		t.Errorf("expected watches to be empty, got %v", w.watches)
	}

	if len(w.events) != 0 {
		t.Errorf("expected events to be empty, got %v", w.events)
	}

	if len(w.errors) != 0 {
		t.Errorf("expected errors to be empty, got %v", w.errors)
	}

	if w.closed != true {
		t.Errorf("expected closed to be true, got %v", w.closed)
	}
}
