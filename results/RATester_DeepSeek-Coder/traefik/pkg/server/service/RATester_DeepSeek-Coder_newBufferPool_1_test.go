package service

import (
	"fmt"
	"testing"
)

func TestNewBufferPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pool := newBufferPool()

	buffer := pool.Get()
	if len(buffer) != bufferPoolSize {
		t.Errorf("Expected buffer size to be %d, got %d", bufferPoolSize, len(buffer))
	}

	pool.Put(buffer)

	buffer2 := pool.Get()
	if len(buffer2) != bufferPoolSize {
		t.Errorf("Expected buffer size to be %d, got %d", bufferPoolSize, len(buffer2))
	}

	if &buffer[0] != &buffer2[0] {
		t.Errorf("Expected the same underlying array, got different arrays")
	}
}
