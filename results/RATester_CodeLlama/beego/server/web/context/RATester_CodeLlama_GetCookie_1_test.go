package context

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Input: &BeegoInput{
			data: map[interface{}]interface{}{
				"cookie": map[string]string{
					"key": "value",
				},
			},
		},
	}
	assert.Equal(t, "value", ctx.GetCookie("key"))
}
