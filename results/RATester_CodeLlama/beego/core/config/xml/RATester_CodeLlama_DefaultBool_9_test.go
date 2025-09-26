package xml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultBool_9(t *testing.T) {
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
	assert.Equal(t, true, c.DefaultBool("key", false))
	assert.Equal(t, true, c.DefaultBool("key1", false))
}
