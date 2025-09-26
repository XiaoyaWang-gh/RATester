package binding

import (
	"fmt"
	"testing"
)

func TestMapFormWithTag_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}
	var testStruct TestStruct
	form := map[string][]string{
		"name": {"test"},
		"age":  {"18"},
	}
	err := MapFormWithTag(&testStruct, form, "form")
	if err != nil {
		t.Errorf("MapFormWithTag error: %v", err)
	}
	if testStruct.Name != "test" {
		t.Errorf("MapFormWithTag error: %v", testStruct.Name)
	}
	if testStruct.Age != 18 {
		t.Errorf("MapFormWithTag error: %v", testStruct.Age)
	}
}
