package validation

import (
	"fmt"
	"testing"
)

func TestSetError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	fieldName := "testField"
	errMsg := "test error message"
	err := v.SetError(fieldName, errMsg)

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if err.Key != fieldName {
		t.Errorf("Expected key to be %s, but got %s", fieldName, err.Key)
	}

	if err.Field != fieldName {
		t.Errorf("Expected field to be %s, but got %s", fieldName, err.Field)
	}

	if err.Tmpl != errMsg {
		t.Errorf("Expected template to be %s, but got %s", errMsg, err.Tmpl)
	}

	if err.Message != errMsg {
		t.Errorf("Expected message to be %s, but got %s", errMsg, err.Message)
	}
}
