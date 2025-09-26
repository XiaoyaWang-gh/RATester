package highlight

import (
	"fmt"
	"testing"
)

func TestWrite_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &byteCountFlexiWriter{}
	p := []byte("hello")
	n, err := w.Write(p)
	if err != nil {
		t.Errorf("Write returned unexpected error: %v", err)
	}
	if n != len(p) {
		t.Errorf("Write returned %d, want %d", n, len(p))
	}
	if w.counter != n {
		t.Errorf("counter = %d, want %d", w.counter, n)
	}
}
