package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	parser := stringParser{}
	value := "test"
	toType := reflect.TypeOf("")

	result, err := parser.parse(value, toType)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := value
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
