package redirect

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestServeHTTP_30(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	m := &moveHandler{
		location:  &url.URL{},
		permanent: true,
	}
	m.ServeHTTP(rw, req)
	if rw.Code != http.StatusMovedPermanently {
		t.Errorf("expected status code %d, got %d", http.StatusMovedPermanently, rw.Code)
	}
	if rw.Header().Get("Location") != m.location.String() {
		t.Errorf("expected location %s, got %s", m.location.String(), rw.Header().Get("Location"))
	}
}
