package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestFile_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, c.File("testdata/test.txt"))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "text/plain; charset=utf-8", rec.Header().Get(HeaderContentType))
	assert.Equal(t, "attachment; filename=\"test.txt\"", rec.Header().Get(HeaderContentDisposition))
	assert.Equal(t, "Hello, World!", rec.Body.String())
}
