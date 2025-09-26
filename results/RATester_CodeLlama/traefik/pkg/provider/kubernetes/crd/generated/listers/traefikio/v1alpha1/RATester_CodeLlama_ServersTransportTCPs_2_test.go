package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServersTransportTCPs_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &serversTransportTCPLister{}
	namespace := "test"

	// when
	result := s.ServersTransportTCPs(namespace)

	// then
	assert.NotNil(t, result)
}
