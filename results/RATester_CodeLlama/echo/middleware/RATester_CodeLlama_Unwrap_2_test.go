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

	e := &ErrKeyAuthMissing{
		Err: errors.New("error"),
	}
	if got := e.Unwrap(); got != e.Err {
		t.Errorf("Unwrap() = %v, want %v", got, e.Err)
	}
}
