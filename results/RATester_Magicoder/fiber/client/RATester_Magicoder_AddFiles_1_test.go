package client

import (
	"fmt"
	"testing"
)

func TestAddFiles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{
		files: []*File{},
	}

	file1 := &File{
		name: "file1",
	}
	file2 := &File{
		name: "file2",
	}

	r.AddFiles(file1, file2)

	if len(r.files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(r.files))
	}

	if r.files[0] != file1 || r.files[1] != file2 {
		t.Errorf("Expected files to be added in the correct order")
	}
}
