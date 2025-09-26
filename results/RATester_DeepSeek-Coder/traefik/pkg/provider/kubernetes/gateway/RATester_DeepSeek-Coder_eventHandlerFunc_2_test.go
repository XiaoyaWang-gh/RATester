package gateway

import (
	"fmt"
	"testing"
	"time"
)

func TestEventHandlerFunc_2(t *testing.T) {
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
	case e := <-events:
		if e != obj {
			t.Errorf("Expected %v, got %v", obj, e)
		}
	case <-time.After(1 * time.Second):
		t.Errorf("Timeout waiting for event")
	}
}
