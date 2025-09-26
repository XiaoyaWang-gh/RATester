package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestBind_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	type testStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testObj := &testStruct{}

	err := ctrl.Bind(testObj)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	if testObj.Field1 != "testValue" {
		t.Errorf("Expected Field1 to be 'testValue', but got '%s'", testObj.Field1)
	}

	if testObj.Field2 != 123 {
		t.Errorf("Expected Field2 to be 123, but got '%d'", testObj.Field2)
	}
}
