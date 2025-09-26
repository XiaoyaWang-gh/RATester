package xml

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestString_33(t *testing.T) {
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
	v, err := c.String("key")
	assert.NoError(t, err)
	assert.Equal(t, "value", v)
}
