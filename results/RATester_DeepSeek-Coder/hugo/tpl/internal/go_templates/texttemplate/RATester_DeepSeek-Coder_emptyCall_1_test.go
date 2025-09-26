package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmptyCall_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fn := reflect.ValueOf(emptyCall)
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}

	result := fn.Call(args)

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
}
