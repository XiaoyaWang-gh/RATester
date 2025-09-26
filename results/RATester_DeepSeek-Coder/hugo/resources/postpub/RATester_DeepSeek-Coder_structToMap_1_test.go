package postpub

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructToMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string
		Field2 int
	}

	testStruct := TestStruct{
		Field1: "test",
		Field2: 123,
	}

	expectedMap := map[string]any{
		"Field1": "test",
		"Field2": 123,
	}

	resultMap := structToMap(testStruct)

	if !reflect.DeepEqual(resultMap, expectedMap) {
		t.Errorf("Expected %v, got %v", expectedMap, resultMap)
	}
}
