package publisher

import (
	"fmt"
	"testing"
)

func Testbackup_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &htmlElementsCollectorWriter{
		pos:   10,
		input: []byte("Hello, World!"),
	}
	w.backup()
	if w.pos != 9 {
		t.Errorf("Expected pos to be 9, but got %d", w.pos)
	}
	if w.r != ',' {
		t.Errorf("Expected r to be ',', but got %c", w.r)
	}
}
