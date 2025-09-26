package http

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestExpectNParameters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	fn := func(tree *matchersTree, s ...string) error {
		return nil
	}
	n := []int{1, 2, 3}

	// when
	result := expectNParameters(fn, n...)

	// then
	assert.NotNil(t, result)
}
