package pageparser

import (
	"fmt"
	"testing"
)

func TestErrorf_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("test input"),
	}

	expectedErr := fmt.Errorf("test error")
	expectedItem := Item{
		Type: tError,
		Err:  expectedErr,
		low:  0,
		high: 10,
	}

	result := l.errorf("test error")

	if result != nil {
		t.Errorf("Expected result to be nil, got %v", result)
	}

	if len(l.items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(l.items))
	}

	item := l.items[0]
	if item.Type != expectedItem.Type {
		t.Errorf("Expected item type to be %v, got %v", expectedItem.Type, item.Type)
	}
	if item.Err.Error() != expectedErr.Error() {
		t.Errorf("Expected item error to be %v, got %v", expectedErr, item.Err)
	}
	if item.low != expectedItem.low {
		t.Errorf("Expected item low to be %v, got %v", expectedItem.low, item.low)
	}
	if item.high != expectedItem.high {
		t.Errorf("Expected item high to be %v, got %v", expectedItem.high, item.high)
	}
}
