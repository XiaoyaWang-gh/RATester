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

	ev := make(chan interface{})
	reh := &resourceEventHandler{ev: ev}

	obj := "test"
	go func() {
		reh.OnDelete(obj)
	}()

	select {
	case received := <-ev:
		if received != obj {
			t.Errorf("Expected %v, got %v", obj, received)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for event")
	}
}
