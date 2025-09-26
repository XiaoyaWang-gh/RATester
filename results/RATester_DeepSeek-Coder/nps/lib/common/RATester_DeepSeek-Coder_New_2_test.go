package common

import (
	"fmt"
	"testing"
)

func TestNew_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pool := &copyBufferPool{}
	pool.New()

	buffer := pool.Get()
	if len(buffer) != PoolSizeCopy {
		t.Errorf("Expected buffer size to be %d, got %d", PoolSizeCopy, len(buffer))
	}

	pool.Put(buffer)
}
