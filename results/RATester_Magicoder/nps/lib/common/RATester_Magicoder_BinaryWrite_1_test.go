package common

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBinaryWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	raw := &bytes.Buffer{}
	BinaryWrite(raw, "test")

	expected := []byte{0, 0, 0, 4, 116, 101, 115, 116}
	if !bytes.Equal(raw.Bytes(), expected) {
		t.Errorf("Expected %v, got %v", expected, raw.Bytes())
	}
}
