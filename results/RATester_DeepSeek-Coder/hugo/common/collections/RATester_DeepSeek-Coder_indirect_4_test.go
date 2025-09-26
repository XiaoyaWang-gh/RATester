package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndirect_4(t *testing.T) {
	type testStruct struct {
		field string
	}

	t.Run("Test with a pointer", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		test := &testStruct{field: "test"}
		rv, isNil := indirect(reflect.ValueOf(test))
		if rv.Kind() != reflect.Ptr || isNil {
			t.Errorf("Expected a pointer and not nil")
		}
	})

	t.Run("Test with a nil pointer", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var test *testStruct
		rv, isNil := indirect(reflect.ValueOf(test))
		if rv.Kind() != reflect.Ptr || !isNil {
			t.Errorf("Expected a nil pointer")
		}
	})

	t.Run("Test with an interface", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		test := testStruct{field: "test"}
		rv, isNil := indirect(reflect.ValueOf(&test))
		if rv.Kind() != reflect.Interface || isNil {
			t.Errorf("Expected an interface and not nil")
		}
	})

	t.Run("Test with a nil interface", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var test interface{}
		rv, isNil := indirect(reflect.ValueOf(test))
		if rv.Kind() != reflect.Interface || !isNil {
			t.Errorf("Expected a nil interface")
		}
	})
}
