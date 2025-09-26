package framework

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpecifiedHTTPBodyHandler_1(t *testing.T) {
	t.Run("should return the specified body", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		body := []byte("test body")
		handler := SpecifiedHTTPBodyHandler(body)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		require.Equal(t, body, rr.Body.Bytes())
	})
}
