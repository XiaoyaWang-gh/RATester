package pageparser

import (
	"fmt"
	"testing"
)

func TestAppend_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.append(Item{Type: ItemType(1)})
	if len(l.items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(l.items))
	}
}
