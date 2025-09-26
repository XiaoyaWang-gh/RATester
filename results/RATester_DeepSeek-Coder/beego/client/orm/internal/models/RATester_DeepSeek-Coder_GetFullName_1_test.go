package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetFullName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct{}
	typ := reflect.TypeOf(testStruct{})
	expected := "main.testStruct"
	result := GetFullName(typ)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
