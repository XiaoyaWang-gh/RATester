package allconfig

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestBaseConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	assert.Equal(t, c.baseConfig, c.BaseConfig())
}
