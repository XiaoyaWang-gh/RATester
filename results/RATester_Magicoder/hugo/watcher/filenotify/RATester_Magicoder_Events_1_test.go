package filenotify

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestEvents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &filePoller{
		events: make(chan fsnotify.Event),
	}

	event := fsnotify.Event{Name: "test_file", Op: fsnotify.Create}
	w.events <- event

	receivedEvent := <-w.Events()

	if receivedEvent.Name != event.Name || receivedEvent.Op != event.Op {
		t.Errorf("Expected event %v, but got %v", event, receivedEvent)
	}
}
