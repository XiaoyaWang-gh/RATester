package json

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDefaultString_6(t *testing.T) {
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
	assert.Equal(t, "value", c.DefaultString("key", "default"))
	assert.Equal(t, "default", c.DefaultString("not_exist_key", "default"))
}
