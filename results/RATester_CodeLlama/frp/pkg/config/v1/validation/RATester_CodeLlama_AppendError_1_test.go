package validation

import (
	"errors"
	"fmt"
	"testing"
)

func TestAppendError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	err := errors.New("error1")
	errs := []error{errors.New("error2"), errors.New("error3")}
	err = AppendError(err, errs...)
	if err.Error() != "error1; error2; error3" {
		t.Errorf("AppendError failed. got %v, want %v", err, "error1; error2; error3")
	}
}
