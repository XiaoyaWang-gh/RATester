package modules

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWrapModuleNotFound_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &collector{}
	err := c.wrapModuleNotFound(errors.New("test"))
	assert.Equal(t, "test: %w", err.Error())
}
