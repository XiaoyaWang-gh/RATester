package tcp

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetHTTPSHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Router{}
	r.httpsHandler = &http.ServeMux{}
	assert.Equal(t, r.httpsHandler, r.GetHTTPSHandler())
}
