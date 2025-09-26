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

	e := ExecError{
		Name: "test",
		Err:  errors.New("test error"),
	}
	if got := e.Unwrap(); got != e.Err {
		t.Errorf("Unwrap() = %v, want %v", got, e.Err)
	}
}
