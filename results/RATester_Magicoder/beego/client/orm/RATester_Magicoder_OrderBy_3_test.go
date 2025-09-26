package orm

import (
	"fmt"
	"testing"
)

func TestOrderBy_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.OrderBy("field1", "field2")

	expected := "ORDER BY field1, field2"
	if qb.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, qb.String())
	}
}
