package k8s

import (
	"fmt"
	"testing"
	"time"
)

func TestOnDelete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ev := make(chan interface{}, 1)
	defer close(ev)

	reh := &ResourceEventHandler{
		Ev: ev,
	}

	obj := "test"

	go reh.OnDelete(obj)

	select {
	case received := <-ev:
		if received != obj {
			t.Errorf("Expected %v, got %v", obj, received)
		}
	case <-time.After(time.Second):
		t.Errorf("Timeout waiting for event")
	}
}
