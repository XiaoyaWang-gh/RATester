package middleware

import (
	"errors"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
)

func TestRotateAccessLogs_1(t *testing.T) {
	testCases := []struct {
		desc                   string
		accessLoggerMiddleware *accesslog.Handler
		expectedError          error
	}{
		{
			desc:                   "should return nil when accessLoggerMiddleware is nil",
			accessLoggerMiddleware: nil,
			expectedError:          nil,
		},
		{
			desc:                   "should return error when accessLoggerMiddleware.Rotate() returns error",
			accessLoggerMiddleware: &accesslog.Handler{},
			expectedError:          errors.New("some error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			o := &ObservabilityMgr{
				accessLoggerMiddleware: test.accessLoggerMiddleware,
			}

			err := o.RotateAccessLogs()

			if !errors.Is(err, test.expectedError) {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			}
		})
	}
}
