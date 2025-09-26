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
	r := strings.NewReader("hello")
	n, err := b.ReadFrom(r)
	if n != 5 {
		t.Errorf("got %d, want 5", n)
	}
	if err != nil {
		t.Errorf("got error %v, want nil", err)
	}
	if b.String() != "hello" {
		t.Errorf("got %q, want %q", b.String(), "hello")
	}
}
