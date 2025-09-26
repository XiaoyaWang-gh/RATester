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

	oracle := &dbBaseOracle{}
	expected := "SELECT TABLE_NAME FROM USER_TABLES"
	actual := oracle.ShowTablesQuery()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
