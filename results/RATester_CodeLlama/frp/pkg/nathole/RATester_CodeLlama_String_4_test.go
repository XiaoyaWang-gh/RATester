package nathole

import (
	"fmt"
	"net"
	"testing"

	"github.com/zeebo/assert"
)

func TestString_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &ChangedAddress{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 1234,
	}
	assert.Equal(t, "127.0.0.1:1234", s.String())
}
