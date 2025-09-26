package orm

import (
	"fmt"
	"testing"
)

func TestSelect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	fields := []string{"field1", "field2"}
	qb.Select(fields...)
	expected := "SELECT field1, field2"
	if qb.String() != expected {
		t.Errorf("Expected %s, got %s", expected, qb.String())
	}
}
