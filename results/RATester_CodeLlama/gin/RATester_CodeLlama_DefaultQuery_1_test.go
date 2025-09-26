package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestDefaultQuery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Params: Params{
			Param{
				Key:   "key",
				Value: "value",
			},
		},
	}
	assert.Equal(t, "value", c.DefaultQuery("key", "default"))
	assert.Equal(t, "default", c.DefaultQuery("key2", "default"))
}
