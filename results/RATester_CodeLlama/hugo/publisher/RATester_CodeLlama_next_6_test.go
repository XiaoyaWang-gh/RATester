package publisher

import (
	"fmt"
	"testing"
)

func TestNext_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &htmlElementsCollectorWriter{}
	l.input = []byte("abc")
	l.pos = 0
	l.width = 0
	l.next()
	if l.r != 'a' {
		t.Errorf("Expected 'a', got %c", l.r)
	}
	if l.width != 1 {
		t.Errorf("Expected width 1, got %d", l.width)
	}
	if l.pos != 1 {
		t.Errorf("Expected pos 1, got %d", l.pos)
	}
}
