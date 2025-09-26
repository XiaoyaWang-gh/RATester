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

	interval := time.Duration(10)
	watcher := NewPollingWatcher(interval)
	if watcher == nil {
		t.Errorf("NewPollingWatcher() = %v, want %v", watcher, nil)
	}
}
