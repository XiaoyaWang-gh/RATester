package template

import (
	"fmt"
	"testing"
	"testing/fstest"
)

func TestreadFileFS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testFS := fstest.MapFS{
		"test.txt": {
			Data: []byte("Hello, World!"),
		},
	}

	readFile := readFileFS(testFS)

	name, b, err := readFile("test.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if name != "test.txt" {
		t.Errorf("Expected name to be 'test.txt', but got '%s'", name)
	}

	if string(b) != "Hello, World!" {
		t.Errorf("Expected content to be 'Hello, World!', but got '%s'", string(b))
	}
}
