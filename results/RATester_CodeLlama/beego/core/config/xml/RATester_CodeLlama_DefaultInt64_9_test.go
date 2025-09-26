package xml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultInt64_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": int64(100),
		},
	}
	assert.Equal(t, int64(100), c.DefaultInt64("key", 0))
	assert.Equal(t, int64(0), c.DefaultInt64("key1", 0))
}
