package k8s

import (
	"fmt"
	"testing"

	v1 "k8s.io/api/core/v1"
)

func TestOnDelete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ev := make(chan interface{})
	reh := &ResourceEventHandler{Ev: ev}

	obj := &v1.Pod{}
	reh.OnDelete(obj)

	select {
	case <-ev:
	default:
		t.Error("Expected event, but got none")
	}
}
