package xml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultInt_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": 1,
		},
	}
	assert.Equal(t, 1, c.DefaultInt("key", 2))
	assert.Equal(t, 2, c.DefaultInt("key1", 2))
}
