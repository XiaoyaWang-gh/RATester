package orm

import (
	"fmt"
	"testing"
)

func TestValues_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	vals := []string{"val1", "val2", "val3"}
	expected := "VALUES (val1, val2, val3)"
	result := qb.Values(vals...)
	if result.String() != expected {
		t.Errorf("Expected %s, but got %s", expected, result.String())
	}
}
