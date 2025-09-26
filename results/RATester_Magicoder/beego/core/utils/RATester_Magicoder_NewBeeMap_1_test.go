package utils

import (
	"fmt"
	"testing"
)

func TestNewBeeMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bm := NewBeeMap()

	if bm == nil {
		t.Error("NewBeeMap() should not return nil")
	}

	if bm.lock == nil {
		t.Error("NewBeeMap().lock should not return nil")
	}

	if bm.bm == nil {
		t.Error("NewBeeMap().bm should not return nil")
	}
}
