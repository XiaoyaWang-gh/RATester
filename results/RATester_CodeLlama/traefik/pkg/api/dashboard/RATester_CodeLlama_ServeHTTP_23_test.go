package dashboard

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServeHTTP_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	g := Handler{
		assets: nil,
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	// when
	g.ServeHTTP(w, r)

	// then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "frame-src 'self' https://traefik.io https://*.traefik.io;", w.Header().Get("Content-Security-Policy"))
	assert.Equal(t, "", w.Header().Get("Content-Type"))
}
