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

	testErr := errors.New("test error")
	he := &HTTPError{
		Code:     http.StatusInternalServerError,
		Message:  "Internal Server Error",
		Internal: nil,
	}
	he = he.WithInternal(testErr)
	if he.Internal != testErr {
		t.Errorf("Expected Internal to be %v, but got %v", testErr, he.Internal)
	}
}
