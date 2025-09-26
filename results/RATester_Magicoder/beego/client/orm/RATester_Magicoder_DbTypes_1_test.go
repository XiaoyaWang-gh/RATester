package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDbTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	oracle := &dbBaseOracle{}
	expected := map[string]string{
		"VARCHAR2": "string",
		"NUMBER":   "int64",
		// add more expected types
	}
	result := oracle.DbTypes()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}
