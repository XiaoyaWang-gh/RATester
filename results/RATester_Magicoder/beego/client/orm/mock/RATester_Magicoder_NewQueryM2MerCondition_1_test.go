package mock

import (
	"fmt"
	"testing"
)

func TestNewQueryM2MerCondition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tableName := "test_table"
	name := "test_name"

	q := NewQueryM2MerCondition(tableName, name)

	if q.tableName != tableName {
		t.Errorf("Expected tableName to be %s, but got %s", tableName, q.tableName)
	}

	if q.name != name {
		t.Errorf("Expected name to be %s, but got %s", name, q.name)
	}
}
