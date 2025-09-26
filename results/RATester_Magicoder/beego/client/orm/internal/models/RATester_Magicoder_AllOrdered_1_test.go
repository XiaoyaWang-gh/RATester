package models

import (
	"fmt"
	"testing"
)

func TestAllOrdered_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		orders: []string{"table1", "table2", "table3"},
		cache: map[string]*ModelInfo{
			"table1": {Table: "table1"},
			"table2": {Table: "table2"},
			"table3": {Table: "table3"},
		},
	}

	result := mc.AllOrdered()

	if len(result) != len(mc.orders) {
		t.Errorf("Expected length of result to be %d, but got %d", len(mc.orders), len(result))
	}

	for i, table := range mc.orders {
		if result[i].Table != table {
			t.Errorf("Expected table %s at index %d, but got %s", table, i, result[i].Table)
		}
	}
}
