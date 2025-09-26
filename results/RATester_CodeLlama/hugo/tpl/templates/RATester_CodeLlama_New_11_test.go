package templates

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/deps"
)

func TestNew_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	deps := &deps.Deps{}
	ns := New(deps)

	assert.NotNil(t, ns)
	assert.Equal(t, deps, ns.deps)
}
