package ping

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServeHTTP_24(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	handler := &Handler{
		EntryPoint:            "entrypoint",
		ManualRouting:         true,
		TerminatingStatusCode: 200,
	}
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	// when
	handler.ServeHTTP(response, request)

	// then
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "OK", response.Body.String())
}
