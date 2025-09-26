package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestQueryArray_1(t *testing.T) {
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
	assert.Equal(t, []string{"value"}, c.QueryArray("key"))
}
