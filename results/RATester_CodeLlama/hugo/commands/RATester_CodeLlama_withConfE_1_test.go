package commands

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWithConfE_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	c := &hugoBuilder{}
	fn := func(conf *commonConfig) error {
		return nil
	}

	// When
	err := c.withConfE(fn)

	// Then
	assert.NoError(t, err)
}
