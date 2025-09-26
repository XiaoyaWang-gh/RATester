package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	r := NewRouter(e)
	assert.NotNil(t, r)
	assert.NotNil(t, r.tree)
	assert.NotNil(t, r.routes)
	assert.Equal(t, e, r.echo)
}
