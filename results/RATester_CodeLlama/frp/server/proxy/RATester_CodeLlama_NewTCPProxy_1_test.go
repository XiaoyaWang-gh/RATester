package proxy

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewTCPProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	baseProxy := &BaseProxy{
		name: "test",
	}
	proxy := NewTCPProxy(baseProxy)
	assert.NotNil(t, proxy)
}
