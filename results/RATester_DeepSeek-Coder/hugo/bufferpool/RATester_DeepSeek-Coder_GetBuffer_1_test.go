package bufferpool

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestGetBuffer_1(t *testing.T) {
	t.Run("should return a new buffer from the pool", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		buf := GetBuffer()
		if buf == nil {
			t.Errorf("Expected a new buffer, got nil")
		}
	})

	t.Run("should return a buffer with a specific type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		buf := GetBuffer()
		if reflect.TypeOf(buf) != reflect.TypeOf(&bytes.Buffer{}) {
			t.Errorf("Expected buffer type to be *bytes.Buffer, got %T", buf)
		}
	})

	t.Run("should return a buffer with a specific length", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		buf := GetBuffer()
		if buf.Len() != 0 {
			t.Errorf("Expected buffer length to be 0, got %d", buf.Len())
		}
	})
}
