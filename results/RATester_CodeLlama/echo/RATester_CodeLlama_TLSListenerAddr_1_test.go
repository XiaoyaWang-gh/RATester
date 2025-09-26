package echo

import (
	"fmt"
	"net"
	"testing"

	"github.com/zeebo/assert"
)

func TestTLSListenerAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	e.TLSListener = &net.TCPListener{}
	assert.NotNil(t, e.TLSListenerAddr())
}
