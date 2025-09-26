package context

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestBindStruct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string
		Field2 int
	}

	testParams := url.Values{
		"Field1": []string{"test1"},
		"Field2": []string{"123"},
	}

	input := &BeegoInput{
		Context: &Context{},
	}

	result := input.bindStruct(&testParams, "TestStruct", reflect.TypeOf(TestStruct{}))

	if result.Field(0).String() != "test1" {
		t.Errorf("Expected Field1 to be 'test1', got %s", result.Field(0).String())
	}

	if result.Field(1).Int() != 123 {
		t.Errorf("Expected Field2 to be 123, got %d", result.Field(1).Int())
	}
}
