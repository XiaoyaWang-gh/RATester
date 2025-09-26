package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestTraefikServices_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &traefikServiceLister{}
	namespace := "test"

	// when
	result := s.TraefikServices(namespace)

	// then
	assert.NotNil(t, result)
}
