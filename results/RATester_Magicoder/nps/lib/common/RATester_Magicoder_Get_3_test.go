package common

import (
	"fmt"
	"sync"
	"testing"
)

func TestGet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pool := &copyBufferPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, PoolSizeCopy)
			},
		},
	}

	buf := pool.Get()
	if len(buf) != PoolSizeCopy {
		t.Errorf("Expected buffer size %d, got %d", PoolSizeCopy, len(buf))
	}

	pool.Put(buf)
}
