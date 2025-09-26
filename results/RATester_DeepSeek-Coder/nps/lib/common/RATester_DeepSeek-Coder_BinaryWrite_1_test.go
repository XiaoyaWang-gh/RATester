package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinaryWrite_1(t *testing.T) {
	t.Run("TestBinaryWrite", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		raw := &bytes.Buffer{}
		BinaryWrite(raw, "test1", "test2", "test3")

		var length int32
		binary.Read(raw, binary.LittleEndian, &length)
		if int(length) != 18 {
			t.Errorf("Expected length to be 18, got %d", length)
		}

		b := make([]byte, length)
		binary.Read(raw, binary.LittleEndian, &b)
		if string(b) != "test1test2test3" {
			t.Errorf("Expected string to be 'test1test2test3', got %s", string(b))
		}
	})
}
