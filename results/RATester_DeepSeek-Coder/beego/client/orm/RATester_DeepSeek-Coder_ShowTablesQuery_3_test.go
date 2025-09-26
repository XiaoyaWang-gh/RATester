package orm

import (
	"fmt"
	"testing"
)

func TestShowTablesQuery_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseMysql{}
	expected := "SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema = DATABASE()"
	result := db.ShowTablesQuery()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
