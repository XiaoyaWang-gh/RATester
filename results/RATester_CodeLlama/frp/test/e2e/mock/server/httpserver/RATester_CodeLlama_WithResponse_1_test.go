package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	resp := []byte("hello")
	opt := WithResponse(resp)
	s := &Server{}
	opt(s)
	if s.handler == nil {
		t.Error("handler should not be nil")
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	s.handler.ServeHTTP(w, r)
	if w.Body.String() != string(resp) {
		t.Errorf("handler should write %s, but got %s", resp, w.Body.String())
	}
}
