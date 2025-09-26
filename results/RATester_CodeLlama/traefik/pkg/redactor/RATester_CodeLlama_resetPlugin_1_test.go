package redactor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestResetPlugin_1(t *testing.T) {
	t.Parallel()
	t.Run("Testcase1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		field := reflect.ValueOf(struct{}{})
		resetPlugin(field)
		if field.Kind() != reflect.Struct {
			t.Errorf("field.Kind() = %v, want %v", field.Kind(), reflect.Struct)
		}
	})
}
