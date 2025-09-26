package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Registering{}
	route := r.Route("/test")
	assert.NotNil(t, route)
}
