package debug

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/deps"
)

func TestNew_57(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	d := &deps.Deps{}

	// When
	ns := New(d)

	// Then
	assert.NotNil(t, ns)
}
