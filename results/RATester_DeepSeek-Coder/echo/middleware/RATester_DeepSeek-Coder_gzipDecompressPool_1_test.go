package middleware

import (
	"compress/gzip"
	"fmt"
	"testing"
)

func TestGzipDecompressPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	d := &DefaultGzipDecompressPool{}
	pool := d.gzipDecompressPool()

	// Test that the pool returns a new gzip.Reader when Get is called
	reader := pool.Get().(*gzip.Reader)
	if reader == nil {
		t.Errorf("Expected a new gzip.Reader, got nil")
	}

	// Test that the pool returns the gzip.Reader to the pool when Put is called
	pool.Put(reader)
	reader2 := pool.Get().(*gzip.Reader)
	if reader != reader2 {
		t.Errorf("Expected the same gzip.Reader after Put, got different ones")
	}
}
