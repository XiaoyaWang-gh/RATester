package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestParamValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{
		pnames:  []string{"name", "age"},
		pvalues: []string{"john", "32"},
	}

	// when
	values := c.ParamValues()

	// then
	assert.Equal(t, []string{"john", "32"}, values)
}
