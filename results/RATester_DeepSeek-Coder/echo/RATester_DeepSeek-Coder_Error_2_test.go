package echo

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestError_2(t *testing.T) {
	t.Run("Test HTTPError.Error()", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		he := &HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  "Internal Server Error",
			Internal: errors.New("database error"),
		}

		expected := "code=500, message=Internal Server Error, internal=database error"
		actual := he.Error()

		if actual != expected {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	})

	t.Run("Test HTTPError.Error() without internal error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		he := &HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}

		expected := "code=500, message=Internal Server Error"
		actual := he.Error()

		if actual != expected {
			t.Errorf("Expected '%s', got '%s'", expected, actual)
		}
	})
}
