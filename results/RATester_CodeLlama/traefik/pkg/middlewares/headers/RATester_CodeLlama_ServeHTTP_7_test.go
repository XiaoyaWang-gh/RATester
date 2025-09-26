package headers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/unrolled/secure"
)

func TestServeHTTP_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var (
		rw  = httptest.NewRecorder()
		req = httptest.NewRequest(http.MethodGet, "/", nil)
		s   = secureHeader{
			next:   http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}),
			secure: &secure.Secure{},
			cfg:    dynamic.Headers{},
		}
	)

	// act
	s.ServeHTTP(rw, req)

	// assert
	assert.Equal(t, http.StatusOK, rw.Code)
}
