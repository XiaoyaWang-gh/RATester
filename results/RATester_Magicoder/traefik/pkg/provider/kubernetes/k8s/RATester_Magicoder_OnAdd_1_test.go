package k8s

import (
	"fmt"
	"testing"
	"time"
)

func TestOnAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ev := make(chan interface{})
	reh := &ResourceEventHandler{
		Ev: ev,
	}

	obj := "testObj"
	go reh.OnAdd(obj, true)

	select {
	case event := <-ev:
		if event != obj {
			t.Errorf("Expected %v, got %v", obj, event)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for event")
	}
}
