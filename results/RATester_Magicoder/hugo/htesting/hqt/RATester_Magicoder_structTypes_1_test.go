package hqt

import (
	"fmt"
	"reflect"
	"testing"
)

func TeststructTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		field1 string
		field2 int
	}

	testValue := reflect.ValueOf(testStruct{
		field1: "test",
		field2: 1,
	})

	testMap := make(map[reflect.Type]struct{})
	structTypes(testValue, testMap)

	if _, ok := testMap[reflect.TypeOf(testStruct{})]; !ok {
		t.Errorf("Expected testStruct to be in the map")
	}
}
