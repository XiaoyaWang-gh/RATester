package proxy

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewGeneralTCPProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	baseProxy := &BaseProxy{}
	proxy := NewGeneralTCPProxy(baseProxy, nil)
	assert.NotNil(t, proxy)
}
