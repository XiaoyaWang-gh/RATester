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

	pool := bufferPool()
	buf := pool.Get().(*bytes.Buffer)
	defer pool.Put(buf)

	buf.WriteString("Hello, World!")
	if buf.String() != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got '%s'", buf.String())
	}
}
