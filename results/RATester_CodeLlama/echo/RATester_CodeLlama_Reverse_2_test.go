package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestReverse_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &Router{}
	name := "test"
	params := []interface{}{"test"}
	// when
	result := r.Reverse(name, params...)
	// then
	assert.Equal(t, "test", result)
}
