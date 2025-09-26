package orm

import (
	"fmt"
	"testing"
)

func TestGroupBy_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	fields := []string{"field1", "field2"}
	expected := "GROUP BY field1, field2"

	qb.GroupBy(fields...)
	result := qb.String()

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
