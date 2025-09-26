package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMiddlewares_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	v := &version{}

	// when
	result := v.Middlewares()

	// then
	assert.NotNil(t, result)
}
