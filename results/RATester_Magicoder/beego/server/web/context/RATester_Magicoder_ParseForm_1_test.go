package context

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParseForm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	form := url.Values{}
	form.Add("field1", "value1")
	form.Add("field2", "123")

	testStruct := TestStruct{}
	err := ParseForm(form, &testStruct)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if testStruct.Field1 != "value1" {
		t.Errorf("Expected Field1 to be 'value1', but got '%s'", testStruct.Field1)
	}

	if testStruct.Field2 != 123 {
		t.Errorf("Expected Field2 to be 123, but got '%d'", testStruct.Field2)
	}
}
