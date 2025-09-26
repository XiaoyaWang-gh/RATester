package json

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestString_27(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": "value",
		},
	}
	v, err := c.String("key")
	assert.NoError(t, err)
	assert.Equal(t, "value", v)
}
