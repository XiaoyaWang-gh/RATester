package json

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestInt_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": 1,
		},
	}
	v, err := c.Int("key")
	assert.NoError(t, err)
	assert.Equal(t, 1, v)
}
