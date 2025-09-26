package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestStaticFS_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()
	router.StaticFS("/static", http.Dir("./static"))

	tests := []struct {
		route string
		code  int
		body  string
	}{
		{"/static/test.txt", http.StatusOK, "This is a test file"},
		{"/static/nonexistent", http.StatusNotFound, ""},
		{"/static", http.StatusNotFound, ""},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, test.code, w.Code)
		if test.body != "" {
			assert.Contains(t, w.Body.String(), test.body)
		}
	}
}
