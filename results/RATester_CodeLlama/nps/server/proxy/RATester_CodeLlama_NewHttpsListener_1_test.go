package proxy

import (
	"fmt"
	"net"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewHttpsListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	l := &net.TCPListener{}

	// act
	httpsListener := NewHttpsListener(l)

	// assert
	assert.NotNil(t, httpsListener)
	assert.NotNil(t, httpsListener.acceptConn)
	assert.Equal(t, l, httpsListener.parentListener)
}
