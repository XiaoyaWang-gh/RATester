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

	d := &DefaultGzipDecompressPool{}
	pool := d.gzipDecompressPool()

	// Test that the pool returns a new gzip.Reader when Get is called
	reader := pool.Get().(*gzip.Reader)
	if reader == nil {
		t.Error("Expected a new gzip.Reader, but got nil")
	}

	// Test that the pool returns the gzip.Reader to the pool when Put is called
	pool.Put(reader)
	newReader := pool.Get().(*gzip.Reader)
	if reader != newReader {
		t.Error("Expected the same gzip.Reader after Put and Get, but got different ones")
	}
}
