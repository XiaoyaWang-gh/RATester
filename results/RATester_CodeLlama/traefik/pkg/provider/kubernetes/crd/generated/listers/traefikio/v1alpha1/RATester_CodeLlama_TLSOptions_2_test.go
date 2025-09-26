package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestTLSOptions_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &tLSOptionLister{}
	namespace := "test"

	// when
	result := s.TLSOptions(namespace)

	// then
	assert.NotNil(t, result)
}
