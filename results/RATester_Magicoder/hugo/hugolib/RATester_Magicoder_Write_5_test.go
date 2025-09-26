package hugolib

import (
	"fmt"
	"testing"
)

func TestWrite_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &lockingBuffer{}
	data := []byte("Hello, World!")
	n, err := b.Write(data)
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
	if n != len(data) {
		t.Errorf("Write returned wrong number of bytes: %d", n)
	}
	if string(b.Bytes()) != string(data) {
		t.Errorf("Write did not write the correct data: %s", b.Bytes())
	}
}
