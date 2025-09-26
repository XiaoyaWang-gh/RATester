package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsEmptyFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	fields := []fieldWithPrefix{
		{
			fieldInfo: &fieldInfo{
				typ: reflect.TypeOf(int(0)),
			},
			prefix: "a",
		},
		{
			fieldInfo: &fieldInfo{
				typ: reflect.TypeOf(int(0)),
			},
			prefix: "b",
		},
	}
	src := map[string][]string{
		"a": {"1"},
		"b": {"2"},
	}

	// Act
	result := isEmptyFields(fields, src)

	// Assert
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
