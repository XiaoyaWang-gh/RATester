package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetSpec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A test context.
	l := &genericResource{}

	// when:
	result := l.getSpec()

	// then:
	assert.Equal(t, l.spec, result)
}
