package limiter

import (
	"fmt"
	"sync"
	"testing"
)

func Testrelease_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &manager{
		pool: sync.Pool{
			New: func() interface{} {
				return &item{}
			},
		},
	}

	item := manager.acquire()
	item.currHits = 1
	item.prevHits = 1
	item.exp = 1

	manager.release(item)

	if item.currHits != 0 || item.prevHits != 0 || item.exp != 0 {
		t.Error("Expected item to be reset")
	}
}
