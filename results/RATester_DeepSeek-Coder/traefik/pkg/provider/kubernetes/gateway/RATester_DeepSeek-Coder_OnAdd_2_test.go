package gateway

import (
	"fmt"
	"testing"
	"time"
)

func TestOnAdd_2(t *testing.T) {
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
	reh.OnAdd(obj, false)

	select {
	case event := <-ev:
		if event != obj {
			t.Errorf("Expected %v, got %v", obj, event)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for event")
	}
}
