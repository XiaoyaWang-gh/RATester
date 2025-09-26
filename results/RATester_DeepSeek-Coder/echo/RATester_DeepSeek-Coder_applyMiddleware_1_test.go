package echo

import (
	"fmt"
	"testing"
)

func TestApplyMiddleware_1(t *testing.T) {
	t.Run("Test applyMiddleware", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testHandler := func(c Context) error {
			return nil
		}

		testMiddleware := func(next HandlerFunc) HandlerFunc {
			return func(c Context) error {
				return next(c)
			}
		}

		result := applyMiddleware(testHandler, testMiddleware)
		if result == nil {
			t.Errorf("Expected a non-nil result")
		}
	})
}
