package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		A string
		B int
		C struct {
			D string
			E int
		}
	}

	test := TestStruct{
		A: "test",
		B: 123,
		C: struct {
			D string
			E int
		}{
			D: "test",
			E: 456,
		},
	}

	expected := map[string]interface{}{
		"A":   "test",
		"B":   123,
		"C.D": "test",
		"C.E": 456,
	}

	actual := make(map[string]interface{})
	list("", test, actual)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
