package filenotify

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

func TestAdd_3(t *testing.T) {
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

	name := "test"
	err := w.Add(name)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, exists := w.watches[name]; !exists {
		t.Fatalf("watch not added")
	}
}
