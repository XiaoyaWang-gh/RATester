package echo

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestWithInternal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	he := &HTTPError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	err := errors.New("internal error")
	he = he.WithInternal(err)
	if he.Internal != err {
		t.Errorf("WithInternal() = %v, want %v", he.Internal, err)
	}
}
