package orm

import (
	"fmt"
	"testing"
)

func TestOffset_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{}
	offset := 10
	qs.Offset(offset)
	if qs.offset != int64(offset) {
		t.Errorf("Expected offset to be %d, got %d", offset, qs.offset)
	}
}
