package filenotify

import (
	"fmt"
	"testing"
)

func TestNewEventWatcher_1(t *testing.T) {
	t.Run("should return a new FileWatcher", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		watcher, err := NewEventWatcher()
		if err != nil {
			t.Errorf("NewEventWatcher() error = %v, wantErr %v", err, false)
			return
		}
		if watcher == nil {
			t.Errorf("NewEventWatcher() = %v, want not nil", watcher)
		}
	})
}
