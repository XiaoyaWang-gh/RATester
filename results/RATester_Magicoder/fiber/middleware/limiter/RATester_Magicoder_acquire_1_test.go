package limiter

import (
	"fmt"
	"sync"
	"testing"
)

func Testacquire_1(t *testing.T) {
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

	if item == nil {
		t.Error("Expected an item, but got nil")
	}

	manager.release(item)

	item2 := manager.acquire()

	if item2 != item {
		t.Error("Expected the same item after release and acquire")
	}
}
