package client

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestApiStop_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	w := httptest.NewRecorder()
	svr.apiStop(w, nil)
	if w.Code != 200 {
		t.Errorf("apiStop failed, code [%d]", w.Code)
	}
}
