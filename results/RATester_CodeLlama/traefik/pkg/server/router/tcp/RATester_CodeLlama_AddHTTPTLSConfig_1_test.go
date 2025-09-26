package tcp

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAddHTTPTLSConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	sniHost := "sniHost"
	config := &tls.Config{}
	r := &Router{}

	// when
	r.AddHTTPTLSConfig(sniHost, config)

	// then
	assert.Equal(t, config, r.hostHTTPTLSConfig[sniHost])
}
