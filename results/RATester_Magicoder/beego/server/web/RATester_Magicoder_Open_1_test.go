package web

import (
	"fmt"
	"io"
	"testing"
)

func TestOpen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := FileSystem{}
	file, err := fs.Open("test.txt")
	if err != nil {
		t.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	_, err = file.Read(make([]byte, 1024))
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		t.Errorf("Failed to seek file: %v", err)
	}
}
