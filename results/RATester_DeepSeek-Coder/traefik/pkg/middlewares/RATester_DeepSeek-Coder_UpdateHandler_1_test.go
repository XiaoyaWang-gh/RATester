package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestUpdateHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HTTPHandlerSwitcher{
		handler: safe.New(http.NotFoundHandler()),
	}

	newHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	h.UpdateHandler(newHandler)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if got, want := rr.Body.String(), "Hello, world!"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
