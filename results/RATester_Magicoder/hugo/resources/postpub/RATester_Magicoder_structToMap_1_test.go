package postpub

import (
	"fmt"
	"reflect"
	"testing"
)

func TeststructToMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string
		Field2 int
		Field3 bool
	}

	testStruct := TestStruct{
		Field1: "test",
		Field2: 123,
		Field3: true,
	}

	expectedMap := map[string]any{
		"Field1": "test",
		"Field2": 123,
		"Field3": true,
	}

	resultMap := structToMap(testStruct)

	if !reflect.DeepEqual(resultMap, expectedMap) {
		t.Errorf("Expected: %v, Got: %v", expectedMap, resultMap)
	}
}
