package param

import (
	"fmt"
	"testing"
)

func TestDefault_1(t *testing.T) {
	t.Run("should return a MethodParamOption that sets the default value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defaultValue := "test"
		option := Default(defaultValue)
		param := &MethodParam{}
		option(param)
		if param.defaultValue != fmt.Sprint(defaultValue) {
			t.Errorf("Expected default value to be %v, got %v", defaultValue, param.defaultValue)
		}
	})

	t.Run("should return a MethodParamOption that does not set the default value if the input is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var defaultValue interface{} = nil
		option := Default(defaultValue)
		param := &MethodParam{}
		option(param)
		if param.defaultValue != "" {
			t.Errorf("Expected default value to be empty, got %v", param.defaultValue)
		}
	})
}
