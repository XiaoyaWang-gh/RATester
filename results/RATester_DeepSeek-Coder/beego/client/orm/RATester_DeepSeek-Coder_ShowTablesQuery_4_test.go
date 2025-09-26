package orm

import (
	"fmt"
	"testing"
)

func TestShowTablesQuery_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseOracle{}
	expected := "SELECT TABLE_NAME FROM USER_TABLES"
	actual := db.ShowTablesQuery()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
