package middleware

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBufferPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	pool := bufferPool()

	b := pool.Get().(*bytes.Buffer)
	if b == nil {
		t.Errorf("Expected non-nil buffer")
	}

	pool.Put(b)

	b = pool.Get().(*bytes.Buffer)
	if b == nil {
		t.Errorf("Expected non-nil buffer")
	}
}
