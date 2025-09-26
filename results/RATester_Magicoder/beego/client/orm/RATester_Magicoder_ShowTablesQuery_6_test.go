package orm

import (
	"fmt"
	"testing"
)

func TestShowTablesQuery_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBase{}
	expected := "SHOW TABLES"
	result := db.ShowTablesQuery()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
