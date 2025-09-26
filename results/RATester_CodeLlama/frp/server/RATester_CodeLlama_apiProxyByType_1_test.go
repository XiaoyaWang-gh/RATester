package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiProxyByType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/proxy/type", nil)
	svr.apiProxyByType(w, r)
	if w.Code != 200 {
		t.Errorf("Http response code [%d] is not 200", w.Code)
	}
	if w.Body.String() != "" {
		t.Errorf("Http response body [%s] is not empty", w.Body.String())
	}
}
