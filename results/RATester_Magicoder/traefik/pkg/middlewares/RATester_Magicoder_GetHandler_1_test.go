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

	handler := &HTTPHandlerSwitcher{
		handler: &safe.Safe{},
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test Handler"))
	})

	handler.UpdateHandler(testHandler)

	newHandler := handler.GetHandler()

	if newHandler == nil {
		t.Error("Expected a handler, but got nil")
	}
}
