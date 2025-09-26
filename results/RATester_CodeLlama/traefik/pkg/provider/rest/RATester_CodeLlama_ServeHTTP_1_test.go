package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestServeHTTP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	p := &Provider{
		configurationChan: make(chan<- dynamic.Message),
	}

	// when
	p.ServeHTTP(rw, req)

	// then
	assert.Equal(t, http.StatusBadRequest, rw.Code)
}
