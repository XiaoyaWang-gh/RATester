package framework

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestFailf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	assert := assert.New(t)
	assert.Panics(func() {
		Failf("test")
	})
}
