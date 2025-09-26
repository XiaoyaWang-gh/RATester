package yaml

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "true",
		},
	}
	b, err := c.Bool("key")
	assert.NoError(t, err)
	assert.True(t, b)
}
