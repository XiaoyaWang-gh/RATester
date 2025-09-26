package orm

import (
	"fmt"
	"testing"
)

func TestSet_35(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Set("a", "b")
	expected := "SET a, b"
	if qb.String() != expected {
		t.Errorf("Expected %s, got %s", expected, qb.String())
	}
}
