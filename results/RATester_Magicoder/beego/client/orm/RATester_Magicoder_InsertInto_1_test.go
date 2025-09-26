package orm

import (
	"fmt"
	"testing"
)

func TestInsertInto_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.InsertInto("table", "field1", "field2")
	expected := "INSERT INTO table (field1, field2)"
	if qb.String() != expected {
		t.Errorf("Expected %s, got %s", expected, qb.String())
	}
}
