package plugins

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGoPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	c.goPath = "test"
	assert.Equal(t, "test", c.GoPath())
}
