package orm

import (
	"fmt"
	"testing"
)

func TestGroupBy_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.GroupBy("field1", "field2")
	expected := "GROUP BY field1, field2"
	if qb.String() != expected {
		t.Errorf("Expected %s, got %s", expected, qb.String())
	}
}
