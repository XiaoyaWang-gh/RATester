package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestReverse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	r := e.Reverse("name", "param1", "param2")
	assert.Equal(t, "/name/param1/param2", r)
}
