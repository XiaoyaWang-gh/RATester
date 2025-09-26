package bufferpool

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPutBuffer_1(t *testing.T) {
	t.Run("Testing PutBuffer", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		buf := bytes.NewBuffer([]byte("Hello, World!"))
		PutBuffer(buf)
		if buf.Len() != 0 {
			t.Errorf("Expected buffer length to be 0, got %d", buf.Len())
		}
	})
}
