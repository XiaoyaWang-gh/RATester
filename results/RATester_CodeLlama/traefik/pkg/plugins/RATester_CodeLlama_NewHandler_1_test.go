package plugins

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestNewHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var ctx context.Context
	var next http.Handler
	var m YaegiMiddleware
	// act
	_, err := m.NewHandler(ctx, next)
	// assert
	if err != nil {
		t.Errorf("NewHandler() error = %v, want nil", err)
		return
	}
}
