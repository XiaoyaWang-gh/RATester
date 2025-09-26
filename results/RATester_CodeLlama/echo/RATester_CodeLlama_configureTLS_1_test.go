package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestConfigureTLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	e := &Echo{}
	address := "127.0.0.1:8080"
	// when
	e.configureTLS(address)
	// then
	assert.Equal(t, address, e.TLSServer.Addr)
	assert.Equal(t, []string{"h2"}, e.TLSServer.TLSConfig.NextProtos)
}
