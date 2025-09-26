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

	buf := pool.Get()
	if len(buf) != PoolSizeCopy {
		t.Errorf("Expected buffer size %d, got %d", PoolSizeCopy, len(buf))
	}

	pool.Put(buf)
	buf2 := pool.Get()
	if len(buf2) != PoolSizeCopy {
		t.Errorf("Expected buffer size %d, got %d", PoolSizeCopy, len(buf2))
	}

	if &buf[0] != &buf2[0] {
		t.Errorf("Expected same buffer address, got different addresses")
	}
}
