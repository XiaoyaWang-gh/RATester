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
	p := []byte("test")
	n, err := b.Write(p)
	if err != nil {
		t.Errorf("Write() error = %v", err)
		return
	}
	if n != len(p) {
		t.Errorf("Write() n = %v, want %v", n, len(p))
	}
}
