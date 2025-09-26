package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestemptyCall_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fn := reflect.ValueOf(emptyCall)
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}
	result := fn.Call(args)
	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result))
	}
	if result[0].Interface() != nil {
		t.Errorf("Expected nil result, got %v", result[0].Interface())
	}
}
