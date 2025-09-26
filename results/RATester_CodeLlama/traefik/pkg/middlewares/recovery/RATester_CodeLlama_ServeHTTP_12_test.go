package recovery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	re := &recovery{
		next: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}),
	}
	re.ServeHTTP(rw, req)
	if rw.Code != http.StatusOK {
		t.Errorf("rw.Code = %d, want %d", rw.Code, http.StatusOK)
	}
}
