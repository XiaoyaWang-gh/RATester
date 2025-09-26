package template

import (
	"fmt"
	"io/fs"
	"testing"
	"testing/fstest"
	"time"
)

func TestReadFileFS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testFS := fstest.MapFS{
		"test.txt": {
			Data:    []byte("Hello, World!"),
			Mode:    fs.ModePerm,
			ModTime: time.Now(),
		},
	}

	readFile := readFileFS(testFS)

	name, data, err := readFile("test.txt")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if name != "test.txt" {
		t.Errorf("Expected name to be 'test.txt', got %s", name)
	}

	if string(data) != "Hello, World!" {
		t.Errorf("Expected data to be 'Hello, World!' got %s", string(data))
	}
}
