package accesslog

import (
	"fmt"
	"os"
	"testing"
)

func TestWrite_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	n := noopCloser{file}

	_, err = n.Write([]byte("Hello, World!"))
	if err != nil {
		t.Fatal(err)
	}

	file.Seek(0, 0)
	buf := make([]byte, 13)
	_, err = file.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	if string(buf) != "Hello, World!" {
		t.Fatalf("Expected 'Hello, World!', got '%s'", string(buf))
	}
}
