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

	e := &ErrKeyAuthMissing{
		Err: errors.New("error"),
	}
	if got := e.Error(); got != "error" {
		t.Errorf("Error() = %v, want %v", got, "error")
	}
}
