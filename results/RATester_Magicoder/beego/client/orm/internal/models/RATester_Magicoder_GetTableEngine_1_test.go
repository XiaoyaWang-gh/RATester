package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetTableEngine_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		TableEngine func() string
	}

	testVal := testStruct{
		TableEngine: func() string {
			return "InnoDB"
		},
	}

	val := reflect.ValueOf(testVal)
	engine := GetTableEngine(val)

	if engine != "InnoDB" {
		t.Errorf("Expected 'InnoDB', got '%s'", engine)
	}
}
