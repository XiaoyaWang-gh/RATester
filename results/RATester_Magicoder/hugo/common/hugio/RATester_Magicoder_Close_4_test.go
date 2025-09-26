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
	defer pr.Close()
	defer pw.Close()

	c := PipeReadWriteCloser{pr, pw}

	err := c.Close()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
