package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestHTMLBlob_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, c.HTMLBlob(http.StatusOK, []byte("")))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, MIMETextHTMLCharsetUTF8, rec.Header().Get(HeaderContentType))
}
