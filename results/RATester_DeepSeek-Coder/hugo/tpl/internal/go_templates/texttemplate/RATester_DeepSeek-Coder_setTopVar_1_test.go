package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetTopVar_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &state{
		vars: []variable{
			{name: "var1", value: reflect.ValueOf(1)},
			{name: "var2", value: reflect.ValueOf(2)},
			{name: "var3", value: reflect.ValueOf(3)},
		},
	}

	newValue := reflect.ValueOf(4)
	s.setTopVar(1, newValue)

	if s.vars[len(s.vars)-1].value.Interface().(int) != 4 {
		t.Errorf("Expected the new value to be 4, but got %v", s.vars[len(s.vars)-1].value.Interface())
	}
}
