package hugio

import (
	"fmt"
	"testing"
)

func TestNewPipeReadWriteCloser_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	prwc := NewPipeReadWriteCloser()

	// Test WriteString
	testString := "Hello, World!"
	n, err := prwc.WriteString(testString)
	if err != nil {
		t.Errorf("WriteString failed: %v", err)
	}
	if n != len(testString) {
		t.Errorf("WriteString wrote %d bytes, expected %d", n, len(testString))
	}

	// Test Read
	buf := make([]byte, len(testString))
	n, err = prwc.Read(buf)
	if err != nil {
		t.Errorf("Read failed: %v", err)
	}
	if n != len(testString) {
		t.Errorf("Read read %d bytes, expected %d", n, len(testString))
	}
	if string(buf) != testString {
		t.Errorf("Read read %q, expected %q", string(buf), testString)
	}

	// Test Close
	err = prwc.Close()
	if err != nil {
		t.Errorf("Close failed: %v", err)
	}
}
