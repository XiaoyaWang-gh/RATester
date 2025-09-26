package gin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBindYAML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// given
	c := &Context{
		Params: Params{
			Param{
				Key:   "key",
				Value: "value",
			},
		},
	}

	// when
	err := c.ShouldBindYAML(nil)

	// then
	assert.NoError(t, err)
}
