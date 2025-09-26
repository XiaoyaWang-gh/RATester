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

	fp := &filePoller{
		events: make(chan fsnotify.Event),
	}

	event := fsnotify.Event{
		Name: "test.txt",
		Op:   fsnotify.Create,
	}

	go func() {
		fp.events <- event
	}()

	receivedEvent := <-fp.Events()

	if receivedEvent.Name != event.Name || receivedEvent.Op != event.Op {
		t.Errorf("Expected event %v, got %v", event, receivedEvent)
	}
}
