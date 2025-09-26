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

	db := &dbBaseOracle{}
	expected := map[string]string{
		"int":     "NUMBER",
		"int8":    "NUMBER",
		"int16":   "NUMBER",
		"int32":   "NUMBER",
		"int64":   "NUMBER",
		"uint":    "NUMBER",
		"uint8":   "NUMBER",
		"uint16":  "NUMBER",
		"uint32":  "NUMBER",
		"uint64":  "NUMBER",
		"float32": "FLOAT",
		"float64": "FLOAT",
		"bool":    "CHAR(1)",
		"string":  "VARCHAR2",
		"time":    "TIMESTAMP",
		"bytes":   "RAW",
	}
	actual := db.DbTypes()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
