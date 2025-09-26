package json

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDefaultFloat_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": float64(1),
		},
	}
	assert.Equal(t, float64(1), c.DefaultFloat("key", 0))
	assert.Equal(t, float64(0), c.DefaultFloat("key1", 0))
}
