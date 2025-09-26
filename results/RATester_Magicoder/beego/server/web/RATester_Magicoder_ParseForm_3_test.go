package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestParseForm_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testObj := testStruct{}

	ctx := &context.Context{}
	ctx.Input.RequestBody = []byte(`field1=value1&field2=123`)

	controller := &Controller{Ctx: ctx}

	err := controller.ParseForm(&testObj)

	if err != nil {
		t.Errorf("Error parsing form: %v", err)
	}

	if testObj.Field1 != "value1" {
		t.Errorf("Expected Field1 to be 'value1', but got '%s'", testObj.Field1)
	}

	if testObj.Field2 != 123 {
		t.Errorf("Expected Field2 to be 123, but got %d", testObj.Field2)
	}
}
