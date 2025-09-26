package orm

import (
	"fmt"
	"testing"
)

func TestMaxLimit_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBase{}
	expected := uint64(18446744073709551615)
	result := db.MaxLimit()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
