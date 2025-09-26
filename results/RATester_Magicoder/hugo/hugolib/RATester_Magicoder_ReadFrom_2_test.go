package hugolib

import (
	"fmt"
	"strings"
	"testing"
)

func TestReadFrom_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &lockingBuffer{}
	r := strings.NewReader("Hello, World!")
	n, err := b.ReadFrom(r)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if n != int64(len("Hello, World!")) {
		t.Errorf("Expected to read %d bytes, but got %d", len("Hello, World!"), n)
	}
	if b.String() != "Hello, World!" {
		t.Errorf("Expected buffer to contain 'Hello, World!', but got '%s'", b.String())
	}
}
