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
	BinaryWrite(raw, "hello", "world")
	if raw.Len() != 10 {
		t.Errorf("BinaryWrite() = %v, want %v", raw.Len(), 10)
	}
}
