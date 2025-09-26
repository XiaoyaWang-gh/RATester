package pageparser

import (
	"fmt"
	"testing"
)

func TestConsume_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	items := []Item{
		{Type: tText, low: 0, high: 5},
		{Type: tError},
		{Type: tEOF},
	}
	iterator := &Iterator{items: items}

	iterator.Consume(3)

	if iterator.lastPos != 1 {
		t.Errorf("Expected lastPos to be 1, but got %d", iterator.lastPos)
	}
}
