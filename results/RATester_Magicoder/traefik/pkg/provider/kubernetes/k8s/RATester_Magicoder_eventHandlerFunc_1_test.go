package k8s

import (
	"fmt"
	"testing"
	"time"
)

func TestEventHandlerFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	events := make(chan interface{}, 1)
	defer close(events)

	obj := "test"

	eventHandlerFunc(events, obj)

	select {
	case event := <-events:
		if event != obj {
			t.Errorf("Expected %v, got %v", obj, event)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timed out waiting for event")
	}
}
