package binding

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	err := SliceValidationError{
		errors.New("error1"),
		errors.New("error2"),
		errors.New("error3"),
	}

	if err.Error() != "error1\nerror2\nerror3" {
		t.Errorf("Expected error1\nerror2\nerror3, got %s", err.Error())
	}
}
