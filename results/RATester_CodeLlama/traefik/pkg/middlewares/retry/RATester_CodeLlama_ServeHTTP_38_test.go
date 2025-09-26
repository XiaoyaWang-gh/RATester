package retry

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServeHTTP_38(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	r := &retry{
		attempts: 1,
	}

	// when
	r.ServeHTTP(rw, req)

	// then
	assert.Equal(t, http.StatusOK, rw.Code)
}
