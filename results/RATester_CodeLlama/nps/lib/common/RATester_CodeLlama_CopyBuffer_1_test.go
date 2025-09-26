package common

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCopyBuffer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var buf bytes.Buffer
	var src bytes.Buffer
	src.WriteString("hello world")
	written, err := CopyBuffer(&buf, &src)
	if err != nil {
		t.Errorf("CopyBuffer error: %v", err)
	}
	if written != int64(src.Len()) {
		t.Errorf("CopyBuffer written: %d, want: %d", written, src.Len())
	}
	if buf.String() != src.String() {
		t.Errorf("CopyBuffer buf: %s, want: %s", buf.String(), src.String())
	}
}
