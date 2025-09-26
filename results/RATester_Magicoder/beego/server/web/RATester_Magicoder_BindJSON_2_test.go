package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestBindJSON_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testObj := &testStruct{}

	err := ctrl.BindJSON(testObj)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if testObj.Field1 != "" || testObj.Field2 != 0 {
		t.Errorf("Expected fields to be empty, but got %v and %v", testObj.Field1, testObj.Field2)
	}
}
