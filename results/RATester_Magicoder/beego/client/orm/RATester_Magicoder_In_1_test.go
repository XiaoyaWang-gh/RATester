package orm

import (
	"fmt"
	"testing"
)

func TestIn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.In("val1", "val2", "val3")
	expected := "IN (val1, val2, val3)"
	if qb.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, qb.String())
	}
}
