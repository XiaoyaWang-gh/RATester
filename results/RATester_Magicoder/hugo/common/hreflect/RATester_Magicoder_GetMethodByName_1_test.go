package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMethodByName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Name string
	}

	testStruct := TestStruct{Name: "Test"}
	v := reflect.ValueOf(testStruct)

	method := GetMethodByName(v, "Name")

	if method.IsValid() {
		t.Errorf("Expected method to be invalid, but it is valid")
	}
}
