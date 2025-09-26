package client

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestHealthz_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	w := httptest.NewRecorder()
	svr.healthz(w, nil)
	if w.Code != 200 {
		t.Fatalf("expect 200, but got %d", w.Code)
	}
}
