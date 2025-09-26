package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zeebo/assert"
)

func TestMultipartForm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("--foo\r\nContent-Disposition: form-data; name=\"foo\"\r\n\r\nbar\r\n--foo--\r\n"))
	req.Header.Set(HeaderContentType, MIMEMultipartForm+"; boundary=foo")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	form, err := c.MultipartForm()
	assert.NoError(t, err)
	assert.NotNil(t, form)
	assert.Equal(t, "bar", form.Value["foo"][0])
}
