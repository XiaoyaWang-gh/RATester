package middleware

import (
	"fmt"
	"testing"
	"time"
)

func TestToMiddleware_1(t *testing.T) {
	t.Run("should return error when timeout is not set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		config := ContextTimeoutConfig{}
		_, err := config.ToMiddleware()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return middleware when timeout is set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		config := ContextTimeoutConfig{Timeout: 1 * time.Second}
		middleware, err := config.ToMiddleware()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if middleware == nil {
			t.Errorf("expected middleware, got nil")
		}
	})
}
