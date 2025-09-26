package filenotify

import (
	"fmt"
	"testing"
	"time"
)

func TestNew_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	interval := time.Duration(10)
	watcher, err := New(interval)
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
	if watcher == nil {
		t.Errorf("New() watcher = nil")
		return
	}
}
