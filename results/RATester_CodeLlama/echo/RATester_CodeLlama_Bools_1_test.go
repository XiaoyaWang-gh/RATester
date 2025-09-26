package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestBools_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A test case
	sourceParam := "sourceParam"
	dest := &[]bool{}
	// and: The method to test
	b := &ValueBinder{}
	b.Bools(sourceParam, dest)
	// when: Invoking the method
	result := b.BindError()
	// then: The result must be nil
	assert.Nil(t, result)
}
