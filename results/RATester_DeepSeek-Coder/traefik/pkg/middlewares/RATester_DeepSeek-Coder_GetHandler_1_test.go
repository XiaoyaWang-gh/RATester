package middlewares

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestGetHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HTTPHandlerSwitcher{
		handler: safe.New(http.NotFoundHandler()),
	}

	newHandler := h.GetHandler()

	if newHandler == nil {
		t.Errorf("Expected a non-nil handler, got nil")
	}
}
