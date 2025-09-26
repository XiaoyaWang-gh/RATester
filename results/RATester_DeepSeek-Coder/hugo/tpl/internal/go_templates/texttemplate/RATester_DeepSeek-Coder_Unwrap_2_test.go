package template

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrap_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testErr := errors.New("test error")
	execErr := ExecError{
		Name: "testTemplate",
		Err:  testErr,
	}

	unwrappedErr := execErr.Unwrap()

	if unwrappedErr != testErr {
		t.Errorf("Expected %v, got %v", testErr, unwrappedErr)
	}
}
