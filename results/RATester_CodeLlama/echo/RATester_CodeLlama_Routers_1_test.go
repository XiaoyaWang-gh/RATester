package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestRouters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	assert.NotNil(t, e.Routers())
}
