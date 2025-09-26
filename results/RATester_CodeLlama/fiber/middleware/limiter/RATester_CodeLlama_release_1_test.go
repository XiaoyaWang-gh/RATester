package limiter

import (
	"fmt"
	"sync"
	"testing"
)

func TestRelease_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	m := &manager{
		pool: sync.Pool{
			New: func() interface{} {
				return &item{}
			},
		},
	}
	e := m.pool.Get().(*item)
	e.prevHits = 1
	e.currHits = 1
	e.exp = 1
	// when
	m.release(e)
	// then
	if e.prevHits != 0 {
		t.Errorf("prevHits should be 0, but got %d", e.prevHits)
	}
	if e.currHits != 0 {
		t.Errorf("currHits should be 0, but got %d", e.currHits)
	}
	if e.exp != 0 {
		t.Errorf("exp should be 0, but got %d", e.exp)
	}
	if m.pool.Get() != e {
		t.Errorf("pool should contain e, but got %v", m.pool.Get())
	}
}
