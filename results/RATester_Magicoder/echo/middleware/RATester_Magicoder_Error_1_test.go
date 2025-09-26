package middleware

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

	testErr := errors.New("test error")
	err := &ErrKeyAuthMissing{Err: testErr}

	if err.Error() != testErr.Error() {
		t.Errorf("Error() = %v, want %v", err.Error(), testErr.Error())
	}
}
