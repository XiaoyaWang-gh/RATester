package auth

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestRequireAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	a := &BasicAuth{
		Secrets: nil,
		Realm:   "test",
	}

	// act
	a.RequireAuth(w, r)

	// assert
	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "Basic realm=\"test\"", w.Header().Get("WWW-Authenticate"))
	assert.Equal(t, "401 Unauthorized\n", w.Body.String())
}
