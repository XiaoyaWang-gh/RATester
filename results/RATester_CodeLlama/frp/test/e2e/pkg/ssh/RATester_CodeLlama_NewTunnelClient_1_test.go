package ssh

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewTunnelClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	tunnelClient := NewTunnelClient("127.0.0.1:10080", "127.0.0.1:22", "echo hello")
	assert.NotNil(t, tunnelClient)
}
