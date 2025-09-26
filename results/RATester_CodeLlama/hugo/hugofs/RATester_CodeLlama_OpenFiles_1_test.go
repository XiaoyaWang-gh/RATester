package hugofs

import (
	"fmt"
	"testing"
)

func TestOpenFiles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var (
		fs = &OpenFilesFs{
			openFiles: map[string]int{},
		}
	)

	// Open a file
	f, err := fs.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Check that the file is open
	if _, ok := fs.OpenFiles()["test.txt"]; !ok {
		t.Error("Expected file to be open")
	}

	// Close the file
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	// Check that the file is no longer open
	if _, ok := fs.OpenFiles()["test.txt"]; ok {
		t.Error("Expected file to be closed")
	}
}
