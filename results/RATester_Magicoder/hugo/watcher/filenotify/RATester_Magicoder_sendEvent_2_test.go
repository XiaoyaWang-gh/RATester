package filenotify

import (
	"fmt"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

func TestsendEvent_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &filePoller{
		events: make(chan fsnotify.Event),
		done:   make(chan struct{}),
	}

	e := fsnotify.Event{Name: "test.txt", Op: fsnotify.Create}

	err := w.sendEvent(e)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	select {
	case event := <-w.events:
		if event.Name != "test.txt" || event.Op != fsnotify.Create {
			t.Errorf("Expected event %v, got %v", e, event)
		}
	case <-time.After(time.Second):
		t.Error("Expected event, got nothing")
	}

	close(w.done)

	err = w.sendEvent(e)
	if err == nil || err.Error() != "closed" {
		t.Errorf("Expected error 'closed', got %v", err)
	}
}
