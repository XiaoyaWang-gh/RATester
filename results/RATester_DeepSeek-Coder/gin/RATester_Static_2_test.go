package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestStatic_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()
	router.Static("/static", "./testdata/dir")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/static/test.txt", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "test data", w.Body.String())
}
