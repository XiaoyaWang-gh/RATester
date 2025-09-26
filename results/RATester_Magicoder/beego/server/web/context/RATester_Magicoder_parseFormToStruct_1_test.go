package context

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestparseFormToStruct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}

	form := url.Values{}
	form.Add("name", "John")
	form.Add("age", "30")

	obj := testStruct{}
	err := parseFormToStruct(form, reflect.TypeOf(obj), reflect.ValueOf(&obj).Elem())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if obj.Name != "John" {
		t.Errorf("Expected Name to be 'John', but got '%s'", obj.Name)
	}

	if obj.Age != 30 {
		t.Errorf("Expected Age to be 30, but got %d", obj.Age)
	}
}
