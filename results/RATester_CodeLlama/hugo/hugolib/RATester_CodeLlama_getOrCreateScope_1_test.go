package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetOrCreateScope_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &cachedContent{}
	pco := &pageContentOutput{}
	scope := "scope"
	cs := c.getOrCreateScope(scope, pco)
	assert.NotNil(t, cs)
	assert.Equal(t, c, cs.cachedContent)
	assert.Equal(t, pco, cs.pco)
	assert.Equal(t, scope, cs.scope)
}
