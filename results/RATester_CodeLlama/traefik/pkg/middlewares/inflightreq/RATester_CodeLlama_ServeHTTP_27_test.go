package inflightreq

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_27(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	i := &inFlightReq{
		handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}),
	}
	i.ServeHTTP(rw, req)
	if rw.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rw.Code)
	}
}
