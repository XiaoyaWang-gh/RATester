package service

import (
	"fmt"
	"testing"
)

func TestnewBufferPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pool := newBufferPool()

	buf := pool.Get()
	if len(buf) != bufferPoolSize {
		t.Errorf("Expected buffer size %d, got %d", bufferPoolSize, len(buf))
	}

	pool.Put(buf)

	buf2 := pool.Get()
	if len(buf2) != bufferPoolSize {
		t.Errorf("Expected buffer size %d, got %d", bufferPoolSize, len(buf2))
	}

	if &buf[0] != &buf2[0] {
		t.Errorf("Expected same buffer address, got different addresses")
	}
}
