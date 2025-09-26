package mirror

import (
	"fmt"
	"sync"
	"testing"
)

func TestInc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Mirroring{
		lock:  sync.RWMutex{},
		total: 0,
	}

	expected := uint64(1)
	result := m.inc()

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
