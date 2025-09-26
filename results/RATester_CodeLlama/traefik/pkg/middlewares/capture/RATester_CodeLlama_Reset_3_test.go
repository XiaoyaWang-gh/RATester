package capture

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestReset_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Capture{}
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("OK"))
	})

	// when
	handler := c.Reset(next)
	handler.ServeHTTP(nil, nil)

	// then
	assert.Equal(t, http.StatusOK, c.StatusCode())
	assert.Equal(t, int64(2), c.RequestSize())
	assert.Equal(t, int64(2), c.ResponseSize())
}
