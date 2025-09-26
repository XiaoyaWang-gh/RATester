package yaml

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDIY_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "value",
		},
	}
	v, err := c.DIY("key")
	assert.NoError(t, err)
	assert.Equal(t, "value", v)
}
