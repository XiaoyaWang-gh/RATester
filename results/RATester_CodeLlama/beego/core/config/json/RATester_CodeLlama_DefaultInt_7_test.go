package json

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDefaultInt_7(t *testing.T) {
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
	assert.Equal(t, 1, c.DefaultInt("key", 2))
	assert.Equal(t, 2, c.DefaultInt("key1", 2))
}
