package middleware

import (
	"compress/gzip"
	"fmt"
	"testing"
)

func TestGzipCompressPool_1(t *testing.T) {
	t.Run("Test gzipCompressPool function", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		config := GzipConfig{
			Level: 5,
		}
		pool := gzipCompressPool(config)
		if pool.New == nil {
			t.Errorf("Expected New to be set")
		}
		w := pool.Get().(*gzip.Writer)
		if w == nil {
			t.Errorf("Expected gzip.Writer, got nil")
		}
		defer w.Close()
	})
}
