package allconfig

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestCreateTitle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	s := "test"
	assert.Equal(t, c.CreateTitle(s), c.config.C.CreateTitle(s))
}
