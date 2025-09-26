package orm

import (
	"fmt"
	"testing"
)

func TestMaxLimit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseSqlite{}
	expected := uint64(9223372036854775807)
	result := db.MaxLimit()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
