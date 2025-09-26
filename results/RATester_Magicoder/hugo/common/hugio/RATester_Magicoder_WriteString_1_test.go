package hugio

import (
	"fmt"
	"io"
	"testing"
)

func TestWriteString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pr, pw := io.Pipe()
	defer pr.Close()
	defer pw.Close()

	c := PipeReadWriteCloser{pr, pw}

	testString := "Hello, World!"
	n, err := c.WriteString(testString)
	if err != nil {
		t.Errorf("WriteString() error = %v", err)
		return
	}

	buf := make([]byte, n)
	_, err = pr.Read(buf)
	if err != nil {
		t.Errorf("Read() error = %v", err)
		return
	}

	if string(buf) != testString {
		t.Errorf("WriteString() = %v, want %v", string(buf), testString)
	}
}
