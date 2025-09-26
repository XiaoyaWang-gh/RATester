package wrr

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServeHTTP_11(t *testing.T) {
	t.Run("should serve the request", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		b := &Balancer{
			stickyCookie: &stickyCookie{
				name: "test",
			},
		}
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		// when
		b.ServeHTTP(w, req)

		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
