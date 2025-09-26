package param

import (
	"fmt"
	"reflect"
	"testing"
)

func Testparse_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	parser := stringParser{}
	value := "test"
	toType := reflect.TypeOf(value)
	result, err := parser.parse(value, toType)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != value {
		t.Errorf("Expected %v, but got %v", value, result)
	}
}
