package orm

import (
	"fmt"
	"testing"
)

func TestShowTablesQuery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseTidb{}
	expected := "SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema = DATABASE()"
	result := db.ShowTablesQuery()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
