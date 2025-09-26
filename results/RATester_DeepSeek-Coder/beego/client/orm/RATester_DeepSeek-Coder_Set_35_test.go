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
	kv := []string{"name", "John", "age", "30"}
	expected := "SET name = John, age = 30"

	qb.Set(kv...)

	if qb.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, qb.String())
	}
}
