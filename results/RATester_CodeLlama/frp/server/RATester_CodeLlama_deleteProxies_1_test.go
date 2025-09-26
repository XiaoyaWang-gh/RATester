package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteProxies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/deleteProxies?status=offline", nil)
	svr.deleteProxies(w, r)
	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
