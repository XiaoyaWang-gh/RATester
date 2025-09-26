package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/bridge"
	"github.com/zeebo/assert"
)

func TestNewWebServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bridge := &bridge.Bridge{}
	s := NewWebServer(bridge)
	assert.NotNil(t, s)
}
