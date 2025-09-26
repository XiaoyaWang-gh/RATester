package tplimpl

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestLoadEmbedded_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Setup test
	assert := assert.New(t)

	// TODO: Test

	// TODO: Teardown
	assert.True(true)
}
