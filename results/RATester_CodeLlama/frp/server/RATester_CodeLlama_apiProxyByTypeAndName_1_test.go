package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiProxyByTypeAndName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/proxy/type/name", nil)
	svr.apiProxyByTypeAndName(w, r)
	if w.Code != 200 {
		t.Errorf("Http response code want [200], but get [%d]", w.Code)
	}
}
