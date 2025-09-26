package allconfig

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetConfigSection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	s := "security"
	assert.Equal(t, c.config.Security, c.GetConfigSection(s))
}
