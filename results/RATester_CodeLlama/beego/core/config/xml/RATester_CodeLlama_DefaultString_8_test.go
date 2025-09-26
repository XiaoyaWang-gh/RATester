package xml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultString_8(t *testing.T) {
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
	assert.Equal(t, "value", c.DefaultString("key", "default"))
	assert.Equal(t, "default", c.DefaultString("key1", "default"))
}
