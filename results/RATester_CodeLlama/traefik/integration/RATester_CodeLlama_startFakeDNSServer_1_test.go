package integration

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestStartFakeDNSServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// setup
	traefikIP := "127.0.0.1"
	// exercise
	srv := startFakeDNSServer(traefikIP)
	// verify
	assert.NotNil(t, srv)
	assert.Equal(t, ":5053", srv.Addr)
	assert.Equal(t, "udp", srv.Net)
	assert.NotNil(t, srv.Handler)
}
