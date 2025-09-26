package yaml

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "value1;value2;value3",
		},
	}
	v, err := c.Strings("key")
	assert.NoError(t, err)
	assert.Equal(t, []string{"value1", "value2", "value3"}, v)
}
