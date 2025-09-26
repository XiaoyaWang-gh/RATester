package filenotify

import (
	"fmt"
	"testing"
	"time"
)

func TestNewPollingWatcher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	interval := time.Duration(10) * time.Second
	watcher := NewPollingWatcher(interval)

	if watcher == nil {
		t.Error("NewPollingWatcher returned nil")
	}

	events := watcher.Events()
	errors := watcher.Errors()

	if events == nil {
		t.Error("Events channel is nil")
	}

	if errors == nil {
		t.Error("Errors channel is nil")
	}
}
