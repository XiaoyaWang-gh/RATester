package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetMiddleware_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var h Handler
	rw := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/middlewares/middlewareID", nil)
	if err != nil {
		t.Fatal(err)
	}

	// when
	h.getMiddleware(rw, request)

	// then
	assert.Equal(t, http.StatusOK, rw.Code)
}
