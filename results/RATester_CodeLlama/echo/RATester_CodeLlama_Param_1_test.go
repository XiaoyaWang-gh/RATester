package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{
		pnames:  []string{"name"},
		pvalues: []string{"john"},
	}
	assert.Equal(t, "john", c.Param("name"))
}
