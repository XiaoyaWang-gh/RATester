package allconfig

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsKindEnabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Config{
		C: &ConfigCompiled{
			DisabledKinds: map[string]bool{
				"kind1": true,
				"kind2": true,
			},
		},
	}

	assert.False(t, c.IsKindEnabled("kind1"))
	assert.False(t, c.IsKindEnabled("kind2"))
	assert.True(t, c.IsKindEnabled("kind3"))
}
