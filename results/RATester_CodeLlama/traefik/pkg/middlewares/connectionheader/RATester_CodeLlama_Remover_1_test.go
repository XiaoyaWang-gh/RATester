package connectionheader

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemover_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		t.Log("next handler")
	})
	remover := Remover(next)
	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	remover.ServeHTTP(rw, req)
	if rw.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rw.Code)
	}
	if rw.Body.String() != "next handler" {
		t.Errorf("expected body %q, got %q", "next handler", rw.Body.String())
	}
}
