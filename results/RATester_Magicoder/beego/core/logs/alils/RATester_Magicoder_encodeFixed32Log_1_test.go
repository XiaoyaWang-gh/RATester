package alils

import (
	"fmt"
	"testing"
)

func TestencodeFixed32Log_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := make([]byte, 10)
	offset := 0
	v := uint32(1234567890)
	expectedOffset := 4

	offset = encodeFixed32Log(data, offset, v)

	if offset != expectedOffset {
		t.Errorf("Expected offset %d, but got %d", expectedOffset, offset)
	}

	expectedV := uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
	if v != expectedV {
		t.Errorf("Expected value %d, but got %d", v, expectedV)
	}
}
