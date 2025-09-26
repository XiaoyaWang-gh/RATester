package tracing

import (
	"fmt"
	"testing"

	"go.opentelemetry.io/otel/codes"
)

func TestServerStatus_1(t *testing.T) {
	t.Run("Testing valid HTTP status codes", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		for i := 100; i < 599; i++ {
			code, message := serverStatus(i)
			if i >= 500 {
				if code != codes.Error || message != "" {
					t.Errorf("Expected (%v, %q), got (%v, %q)", codes.Error, "", code, message)
				}
			} else {
				if code != codes.Unset || message != "" {
					t.Errorf("Expected (%v, %q), got (%v, %q)", codes.Unset, "", code, message)
				}
			}
		}
	})

	t.Run("Testing invalid HTTP status codes", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		code, message := serverStatus(0)
		if code != codes.Error || message != "Invalid HTTP status code 0" {
			t.Errorf("Expected (%v, %q), got (%v, %q)", codes.Error, "Invalid HTTP status code 0", code, message)
		}

		code, message = serverStatus(600)
		if code != codes.Error || message != "Invalid HTTP status code 600" {
			t.Errorf("Expected (%v, %q), got (%v, %q)", codes.Error, "Invalid HTTP status code 600", code, message)
		}
	})
}
