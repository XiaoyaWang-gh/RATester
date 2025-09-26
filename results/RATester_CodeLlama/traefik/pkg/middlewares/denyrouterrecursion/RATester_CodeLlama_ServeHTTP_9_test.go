package denyrouterrecursion

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	l := &DenyRouterRecursion{
		routerName:     "test",
		routerNameHash: "test",
		next:           http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}),
	}

	l.ServeHTTP(rw, req)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rw.Code)
	}
}
