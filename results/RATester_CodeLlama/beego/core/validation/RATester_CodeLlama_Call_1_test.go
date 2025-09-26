package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCall_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := Funcs{
		"add": reflect.ValueOf(func(a, b int) int {
			return a + b
		}),
	}
	result, err := f.Call("add", 1, 2)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Errorf("result length is not 1")
	}
	if result[0].Int() != 3 {
		t.Errorf("result is not 3")
	}
}
