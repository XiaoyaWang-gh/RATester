package lazy

import (
	"fmt"
	"sync"
	"testing"
)

func TestResetWithLock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	om := &onceMore{
		mu:   sync.Mutex{},
		lock: 0,
		done: 0,
	}

	mu := om.ResetWithLock()

	if mu == nil {
		t.Error("Expected a non-nil Mutex, got nil")
	}

	if om.done != 0 {
		t.Error("Expected done to be 0, got", om.done)
	}
}
