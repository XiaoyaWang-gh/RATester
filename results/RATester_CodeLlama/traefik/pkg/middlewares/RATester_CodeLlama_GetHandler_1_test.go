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

	h := &HTTPHandlerSwitcher{}
	h.handler = &safe.Safe{}
	h.handler.Set(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))
	newHandler := h.GetHandler()
	if newHandler == nil {
		t.Error("newHandler should not be nil")
	}
}
