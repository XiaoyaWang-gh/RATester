package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConnectHandler_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var (
		rw  = httptest.NewRecorder()
		req = httptest.NewRequest(http.MethodConnect, "http://example.com", nil)
	)

	hp := &HTTPProxy{}
	hp.ConnectHandler(rw, req)

	if rw.Code != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, rw.Code)
	}
}
