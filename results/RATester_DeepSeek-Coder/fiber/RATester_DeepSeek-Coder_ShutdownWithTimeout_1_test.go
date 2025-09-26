package fiber

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

func TestShutdownWithTimeout_1(t *testing.T) {
	app := &App{
		server: &fasthttp.Server{},
	}

	testCases := []struct {
		name     string
		timeout  time.Duration
		expected error
	}{
		{
			name:     "Shutdown with timeout",
			timeout:  1 * time.Second,
			expected: nil,
		},
		{
			name:     "Shutdown with timeout exceeded",
			timeout:  0 * time.Second,
			expected: context.DeadlineExceeded,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := app.ShutdownWithTimeout(tc.timeout)
			if err != tc.expected {
				t.Errorf("Expected error %v, got %v", tc.expected, err)
			}
		})
	}
}
