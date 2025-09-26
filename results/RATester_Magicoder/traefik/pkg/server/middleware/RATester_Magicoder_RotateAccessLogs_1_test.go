package middleware

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
)

func TestRotateAccessLogs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ObservabilityMgr{
		accessLoggerMiddleware: &accesslog.Handler{},
	}

	err := o.RotateAccessLogs()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
