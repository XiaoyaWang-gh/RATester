package pageparser

import (
	"fmt"
	"testing"
)

func TestEmitString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("test"),
		pos:   2,
		start: 1,
	}

	l.emitString(ItemType(1))

	if l.items[0].Type != ItemType(1) {
		t.Errorf("Expected ItemType to be %v, got %v", ItemType(1), l.items[0].Type)
	}

	if l.items[0].low != 1 {
		t.Errorf("Expected low to be 1, got %v", l.items[0].low)
	}

	if l.items[0].high != 2 {
		t.Errorf("Expected high to be 2, got %v", l.items[0].high)
	}

	if !l.items[0].isString {
		t.Errorf("Expected isString to be true, got false")
	}

	if l.start != 2 {
		t.Errorf("Expected start to be 2, got %v", l.start)
	}
}
