package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetTableName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Name string
	}

	testStruct := TestStruct{Name: "Test"}
	val := reflect.ValueOf(testStruct)

	tableName := GetTableName(val)

	if tableName != "test_struct" {
		t.Errorf("Expected table name to be 'test_struct', but got '%s'", tableName)
	}
}
