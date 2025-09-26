package xml

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestStrings_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "a;b;c",
		},
	}
	v, err := c.Strings("key")
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c"}, v)
}
