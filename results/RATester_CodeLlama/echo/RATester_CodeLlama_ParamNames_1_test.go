package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestParamNames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{
		pnames: []string{"name", "age"},
	}

	// when
	result := c.ParamNames()

	// then
	assert.Equal(t, []string{"name", "age"}, result)
}
