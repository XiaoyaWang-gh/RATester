package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestputBuffer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := &bytes.Buffer{}
	putBuffer(buf)
	if buf.Len() != 0 {
		t.Errorf("Expected buffer length to be 0, but got %d", buf.Len())
	}
}
