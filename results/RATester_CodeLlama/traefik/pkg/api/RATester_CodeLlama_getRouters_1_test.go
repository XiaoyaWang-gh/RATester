package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetRouters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var h Handler
	rw := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	// when
	h.getRouters(rw, request)
	// then
	assert.Equal(t, http.StatusOK, rw.Code)
}
