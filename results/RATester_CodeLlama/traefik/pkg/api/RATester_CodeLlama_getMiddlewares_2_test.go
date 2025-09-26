package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMiddlewares_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var h Handler
	rw := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// when
	h.getMiddlewares(rw, request)

	// then
	// TODO: add assertions
}
