package orm

import (
	"fmt"
	"testing"
)

func TestMaxLimit_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBasePostgres{}
	expected := uint64(0)
	result := db.MaxLimit()
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
