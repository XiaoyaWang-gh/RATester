package failover

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_25(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &Failover{
		handler: &http.ServeMux{},
	}
	f.SetHandlerStatus(context.Background(), true)
	f.SetFallbackHandlerStatus(context.Background(), true)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	f.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
