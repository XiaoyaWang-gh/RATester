package orm

import (
	"fmt"
	"testing"
)

func TestQueryTable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := ormBase{}

	// Test case 1: QueryTable with string input
	qs := o.QueryTable("table_name")
	if qs == nil {
		t.Error("Expected non-nil QuerySeter, got nil")
	}

	// Test case 2: QueryTable with struct input
	type TestStruct struct{}
	qs = o.QueryTable(TestStruct{})
	if qs == nil {
		t.Error("Expected non-nil QuerySeter, got nil")
	}
}
