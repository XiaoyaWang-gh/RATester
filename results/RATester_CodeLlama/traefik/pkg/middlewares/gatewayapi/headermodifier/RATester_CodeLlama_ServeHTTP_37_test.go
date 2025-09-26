package headermodifier

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_37(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	r := &responseHeaderModifier{
		next: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}),
		name: "test",
		set: map[string]string{
			"test": "test",
		},
		add: map[string]string{
			"test": "test",
		},
		remove: []string{"test"},
	}
	r.ServeHTTP(rw, req)
	if rw.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rw.Code)
	}
	if rw.Header().Get("test") != "test" {
		t.Errorf("expected header test to be test, got %s", rw.Header().Get("test"))
	}
}
