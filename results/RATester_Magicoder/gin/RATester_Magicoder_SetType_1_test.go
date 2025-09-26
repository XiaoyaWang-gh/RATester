package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testError := &Error{
		Err:  errors.New("test error"),
		Type: 0,
		Meta: nil,
	}

	testType := ErrorType(1)

	testError.SetType(testType)

	if testError.Type != testType {
		t.Errorf("Expected Type to be %v, but got %v", testType, testError.Type)
	}
}
