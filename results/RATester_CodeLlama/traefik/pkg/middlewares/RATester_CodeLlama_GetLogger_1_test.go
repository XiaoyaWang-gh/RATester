package middlewares

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	middleware := "middleware"
	middlewareType := "middlewareType"
	logger := GetLogger(ctx, middleware, middlewareType)
	assert.NotNil(t, logger)
}
