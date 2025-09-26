package orm

import (
	"fmt"
	"testing"
)

func TestDbTypes_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBase{}
	result := db.DbTypes()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
