package middleware

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
	e := ErrKeyAuthMissing{Err: testErr}

	if e.Unwrap() != testErr {
		t.Errorf("Expected Unwrap to return %v, but got %v", testErr, e.Unwrap())
	}
}
