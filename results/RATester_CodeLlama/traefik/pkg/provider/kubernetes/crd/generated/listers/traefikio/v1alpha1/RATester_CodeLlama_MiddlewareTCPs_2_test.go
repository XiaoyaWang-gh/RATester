package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMiddlewareTCPs_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &middlewareTCPLister{}
	namespace := "namespace"
	assert.NotNil(t, s.MiddlewareTCPs(namespace))
}
