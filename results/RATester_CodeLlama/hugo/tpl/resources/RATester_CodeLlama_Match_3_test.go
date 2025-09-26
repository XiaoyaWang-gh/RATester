package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMatch_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	pattern := "foo"
	ns := &Namespace{}

	// When
	r := ns.Match(pattern)

	// Then
	assert.NotNil(t, r)
}
