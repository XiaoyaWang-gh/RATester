package redirect

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_31(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	r := redirect{
		name: "test",
	}
	r.ServeHTTP(rw, req)
	if rw.Code != http.StatusMovedPermanently {
		t.Errorf("expected status code %d, got %d", http.StatusMovedPermanently, rw.Code)
	}
	if rw.Header().Get("Location") != "/" {
		t.Errorf("expected location header %s, got %s", "/", rw.Header().Get("Location"))
	}
}
