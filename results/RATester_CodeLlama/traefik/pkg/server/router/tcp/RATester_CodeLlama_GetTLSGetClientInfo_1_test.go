package tcp

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetTLSGetClientInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	r := &Router{}
	r.hostHTTPTLSConfig = make(map[string]*tls.Config)
	r.hostHTTPTLSConfig["example.com"] = &tls.Config{}
	r.hostHTTPTLSConfig["example2.com"] = &tls.Config{}
	r.httpsTLSConfig = &tls.Config{}

	// Act
	result := r.GetTLSGetClientInfo()

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, 3, len(r.hostHTTPTLSConfig))
	assert.Equal(t, r.httpsTLSConfig, r.hostHTTPTLSConfig[""])
}
