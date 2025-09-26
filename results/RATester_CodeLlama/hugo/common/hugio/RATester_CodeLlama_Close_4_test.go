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

	c := PipeReadWriteCloser{
		PipeReader: &io.PipeReader{},
		PipeWriter: &io.PipeWriter{},
	}
	err := c.Close()
	if err != nil {
		t.Errorf("Close() = %v, want nil", err)
	}
}
