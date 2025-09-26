package gateway

import (
	"fmt"
	"testing"
	"time"
)

func TestOnDelete_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ev := make(chan interface{}, 1)
	defer close(ev)

	reh := &resourceEventHandler{
		ev: ev,
	}

	obj := "test"

	go reh.OnDelete(obj)

	select {
	case received := <-ev:
		if received != obj {
			t.Errorf("Expected %v, got %v", obj, received)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for event")
	}
}
