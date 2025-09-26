package render

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFlush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BufWriter{
		Buffer: &bytes.Buffer{},
	}

	err := b.Flush()
	if err != nil {
		t.Errorf("Flush() returned an error: %v", err)
	}
}
