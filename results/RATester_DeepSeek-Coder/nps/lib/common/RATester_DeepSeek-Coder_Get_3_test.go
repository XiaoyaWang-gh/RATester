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

	buffer := pool.Get()
	if len(buffer) != PoolSizeCopy {
		t.Errorf("Expected buffer length to be %d, got %d", PoolSizeCopy, len(buffer))
	}
}
