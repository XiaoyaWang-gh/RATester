package orm

import (
	"fmt"
	"testing"
)

func TestsetOffset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := &querySet{}
	qs.setOffset(10)
	if qs.offset != 10 {
		t.Errorf("Expected offset to be 10, but got %d", qs.offset)
	}
}
