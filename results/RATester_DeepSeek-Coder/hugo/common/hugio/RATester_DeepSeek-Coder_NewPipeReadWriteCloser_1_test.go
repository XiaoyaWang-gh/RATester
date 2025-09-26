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

	if prwc.PipeReader == nil || prwc.PipeWriter == nil {
		t.Errorf("NewPipeReadWriteCloser() should return a PipeReadWriteCloser with both PipeReader and PipeWriter")
	}

	testData := "Hello, World!"

	go func() {
		_, err := prwc.WriteString(testData)
		if err != nil {
			t.Errorf("WriteString() error = %v", err)
		}
	}()

	buf := make([]byte, len(testData))
	n, err := prwc.Read(buf)
	if err != nil {
		t.Errorf("Read() error = %v", err)
		return
	}

	if n != len(testData) {
		t.Errorf("Read() = %v, want %v", n, len(testData))
	}

	if string(buf) != testData {
		t.Errorf("Read() = %v, want %v", string(buf), testData)
	}

	err = prwc.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}
}
