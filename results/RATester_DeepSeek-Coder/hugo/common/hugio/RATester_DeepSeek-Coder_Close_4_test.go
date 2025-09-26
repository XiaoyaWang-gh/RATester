package hugio

import (
	"fmt"
	"io"
	"testing"
)

func TestClose_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pr, pw := io.Pipe()
	c := PipeReadWriteCloser{PipeReader: pr, PipeWriter: pw}

	err := c.Close()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = pr.Read(make([]byte, 1))
	if err != io.ErrClosedPipe {
		t.Errorf("Expected io.ErrClosedPipe, got %v", err)
	}

	_, err = pw.Write(make([]byte, 1))
	if err != io.ErrClosedPipe {
		t.Errorf("Expected io.ErrClosedPipe, got %v", err)
	}
}
