package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPutBuffer_1(t *testing.T) {
	t.Run("Test putBuffer function", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		buf := bytes.NewBuffer(make([]byte, 0, 1025))
		putBuffer(buf)
		if buf.Len() != 0 {
			t.Errorf("Expected buffer length to be 0, got %d", buf.Len())
		}
	})
}
