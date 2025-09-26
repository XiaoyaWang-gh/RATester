package middleware

import (
	"compress/gzip"
	"fmt"
	"testing"
)

func TestGzipCompressPool_1(t *testing.T) {
	config := GzipConfig{
		Level:     -1,
		MinLength: 0,
	}

	pool := gzipCompressPool(config)

	t.Run("Test gzipCompressPool with valid config", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		w := pool.Get().(*gzip.Writer)
		defer w.Close()

		_, err := w.Write([]byte("test data"))
		if err != nil {
			t.Errorf("Error writing to gzip writer: %v", err)
		}

		err = w.Close()
		if err != nil {
			t.Errorf("Error closing gzip writer: %v", err)
		}
	})

	t.Run("Test gzipCompressPool with invalid config", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		config.Level = 100
		pool := gzipCompressPool(config)

		w := pool.Get().(*gzip.Writer)
		defer w.Close()

		_, err := w.Write([]byte("test data"))
		if err == nil {
			t.Error("Expected error writing to gzip writer, but got nil")
		}

		err = w.Close()
		if err == nil {
			t.Error("Expected error closing gzip writer, but got nil")
		}
	})
}
