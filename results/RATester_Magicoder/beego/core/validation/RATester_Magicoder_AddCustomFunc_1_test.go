package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddCustomFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	funcs = make(map[string]reflect.Value)
	unFuncs = make(map[string]bool)

	testFunc := func(v *Validation, obj interface{}, key string) {
		// test function implementation
	}

	err := AddCustomFunc("testFunc", testFunc)
	if err != nil {
		t.Errorf("AddCustomFunc returned an error: %v", err)
	}

	if _, ok := funcs["testFunc"]; !ok {
		t.Errorf("AddCustomFunc did not add the function to the funcs map")
	}

	err = AddCustomFunc("testFunc", testFunc)
	if err == nil {
		t.Errorf("AddCustomFunc did not return an error for duplicate function name")
	}
}
