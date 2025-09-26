package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}

	v.AddError("field.name", "error message")

	if len(v.Errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(v.Errors))
	}

	expectedError := &Error{
		Message: "field.name error message",
		Key:     "field.name",
		Name:    "name",
		Field:   "field",
		Label:   "field",
	}

	if !reflect.DeepEqual(v.Errors[0], expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, v.Errors[0])
	}
}
