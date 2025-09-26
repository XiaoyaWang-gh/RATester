package orm

import (
	"fmt"
	"testing"
)

func TestDelete_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Delete("table1", "table2")
	expected := "DELETE table1, table2"
	if qb.String() != expected {
		t.Errorf("Expected %s, got %s", expected, qb.String())
	}
}
