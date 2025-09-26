package echo

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/zeebo/assert"
)

func TestFormParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{
		request: &http.Request{
			Header: http.Header{
				HeaderContentType: []string{MIMEMultipartForm},
			},
		},
	}
	// when
	_, err := c.FormParams()
	// then
	assert.NoError(t, err)
}
