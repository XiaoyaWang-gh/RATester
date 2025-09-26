package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrap_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testErr := errors.New("test error")
	he := &HTTPError{
		Internal: testErr,
	}

	unwrappedErr := he.Unwrap()

	if unwrappedErr != testErr {
		t.Errorf("Expected Unwrap to return %v, but got %v", testErr, unwrappedErr)
	}
}
