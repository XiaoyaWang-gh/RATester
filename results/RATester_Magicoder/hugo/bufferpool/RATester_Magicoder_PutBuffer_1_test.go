package bufferpool

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPutBuffer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := &bytes.Buffer{}
	PutBuffer(buf)

	if buf.Len() != 0 {
		t.Errorf("Expected buffer length to be 0, but got %d", buf.Len())
	}

	if buf.Cap() != 0 {
		t.Errorf("Expected buffer capacity to be 0, but got %d", buf.Cap())
	}

	if buf.String() != "" {
		t.Errorf("Expected buffer string to be empty, but got %s", buf.String())
	}
}
