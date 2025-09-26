package observability

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestNewService_1(t *testing.T) {
	t.Run("NewService", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		service := "service"
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
		got := NewService(ctx, service, next)
		if got == nil {
			t.Errorf("NewService() = %v, want %v", got, "not nil")
		}
	})
}
