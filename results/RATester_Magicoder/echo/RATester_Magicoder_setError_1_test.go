package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{}
	err1 := errors.New("error1")
	err2 := errors.New("error2")

	b.setError(err1)
	if len(b.errors) != 1 || b.errors[0] != err1 {
		t.Errorf("Expected error to be set correctly")
	}

	b.setError(err2)
	if len(b.errors) != 2 || b.errors[1] != err2 {
		t.Errorf("Expected error to be set correctly")
	}
}
