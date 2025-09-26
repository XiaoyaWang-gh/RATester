package gateway

import (
	"fmt"
	"testing"

	v1 "k8s.io/api/core/v1"
)

func TestOnUpdate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ev := make(chan interface{})
	reh := &resourceEventHandler{ev: ev}

	newObj := &v1.Pod{}
	reh.OnUpdate(nil, newObj)

	select {
	case <-ev:
	default:
		t.Error("expected event")
	}
}
