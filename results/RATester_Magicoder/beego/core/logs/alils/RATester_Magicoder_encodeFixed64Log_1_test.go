package alils

import (
	"fmt"
	"testing"
)

func TestencodeFixed64Log_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := make([]byte, 10)
	offset := 0
	v := uint64(1234567890)
	expectedOffset := 8

	actualOffset := encodeFixed64Log(data, offset, v)

	if actualOffset != expectedOffset {
		t.Errorf("Expected offset %d, but got %d", expectedOffset, actualOffset)
	}

	expectedData := []byte{0, 0, 0, 0, 0, 0, 0x1E, 0x4E}
	for i := 0; i < 8; i++ {
		if data[i] != expectedData[i] {
			t.Errorf("Expected data[%d] to be %d, but got %d", i, expectedData[i], data[i])
		}
	}
}
