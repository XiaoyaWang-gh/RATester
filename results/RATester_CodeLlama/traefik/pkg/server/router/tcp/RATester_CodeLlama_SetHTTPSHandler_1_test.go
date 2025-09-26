package tcp

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetHTTPSHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	r := &Router{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	config := &tls.Config{}

	// Act
	r.SetHTTPSHandler(handler, config)

	// Assert
	assert.Equal(t, handler, r.httpsHandler)
	assert.Equal(t, config, r.httpsTLSConfig)
}
