package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCall_1(t *testing.T) {
	type testStruct struct {
		name string
		age  int
	}

	testFunc := func(name string, age int) testStruct {
		return testStruct{name: name, age: age}
	}

	funcs := make(Funcs)
	funcs["testFunc"] = reflect.ValueOf(testFunc)

	t.Run("TestCall", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		params := []interface{}{"John", 30}
		result, err := funcs.Call("testFunc", params...)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(result) != 1 {
			t.Errorf("Expected 1 result, got %d", len(result))
		}
		res := result[0].Interface().(testStruct)
		if res.name != "John" || res.age != 30 {
			t.Errorf("Expected testStruct{name: \"John\", age: 30}, got %v", res)
		}
	})

	t.Run("TestCallWithWrongNumberOfParams", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		params := []interface{}{"John"}
		_, err := funcs.Call("testFunc", params...)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("TestCallWithNonExistentFunction", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		params := []interface{}{"John", 30}
		_, err := funcs.Call("nonExistentFunc", params...)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
