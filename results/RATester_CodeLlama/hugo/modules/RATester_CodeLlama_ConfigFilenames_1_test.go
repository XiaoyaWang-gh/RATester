package modules

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestConfigFilenames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{
		configFilenames: []string{"a", "b"},
	}
	assert.Equal(t, []string{"a", "b"}, m.ConfigFilenames())
}
