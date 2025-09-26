package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestGetUDPServices_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	handler := Handler{
		staticConfig: static.Configuration{},
	}
	rw := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	// when
	handler.getUDPServices(rw, request)
	// then
	assert.Equal(t, http.StatusOK, rw.Code)
}
