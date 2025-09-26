package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMiddlewares_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &middlewareLister{}
	namespace := "test"

	// when
	result := s.Middlewares(namespace)

	// then
	assert.NotNil(t, result)
}
