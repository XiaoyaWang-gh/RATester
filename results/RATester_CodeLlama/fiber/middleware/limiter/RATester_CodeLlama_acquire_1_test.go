package limiter

import (
	"fmt"
	"sync"
	"testing"
)

func TestAcquire_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &manager{
		pool: sync.Pool{
			New: func() interface{} {
				return &item{}
			},
		},
	}
	it := m.acquire()
	if it == nil {
		t.Error("acquire() should not return nil")
	}
	if it.currHits != 0 {
		t.Errorf("acquire() should set currHits to 0, but got %d", it.currHits)
	}
	if it.prevHits != 0 {
		t.Errorf("acquire() should set prevHits to 0, but got %d", it.prevHits)
	}
}
