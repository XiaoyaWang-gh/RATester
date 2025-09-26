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

	he := &HTTPError{
		Internal: errors.New("internal error"),
	}
	if got := he.Unwrap(); got != he.Internal {
		t.Errorf("HTTPError.Unwrap() = %v, want %v", got, he.Internal)
	}
}
