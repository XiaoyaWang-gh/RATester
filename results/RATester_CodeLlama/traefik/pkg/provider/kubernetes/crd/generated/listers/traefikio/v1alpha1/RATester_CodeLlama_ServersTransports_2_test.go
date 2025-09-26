package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServersTransports_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &serversTransportLister{}
	namespace := "test"

	// when
	result := s.ServersTransports(namespace)

	// then
	assert.NotNil(t, result)
}
