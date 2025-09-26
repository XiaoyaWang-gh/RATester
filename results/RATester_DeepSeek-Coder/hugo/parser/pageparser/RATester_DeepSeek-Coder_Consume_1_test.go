package pageparser

import (
	"errors"
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
		{Type: tError, Err: errors.New("error")},
		{Type: tText, low: 6, high: 10},
	}
	iter := &Iterator{items: items}

	iter.Consume(3)

	if iter.Pos() != 2 {
		t.Errorf("Expected position 2, got %d", iter.Pos())
	}

	iter.Consume(2)

	if iter.Pos() != 2 {
		t.Errorf("Expected position 2, got %d", iter.Pos())
	}
}
