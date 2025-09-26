package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestBindYAML_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	type testStruct struct {
		Field1 string `yaml:"field1"`
		Field2 int    `yaml:"field2"`
	}

	testObj := &testStruct{}

	err := ctrl.BindYAML(testObj)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if testObj.Field1 != "test" {
		t.Errorf("Expected Field1 to be 'test', but got %s", testObj.Field1)
	}

	if testObj.Field2 != 123 {
		t.Errorf("Expected Field2 to be 123, but got %d", testObj.Field2)
	}
}
