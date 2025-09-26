package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestAcquireContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	c := e.AcquireContext()
	assert.NotNil(t, c)
	assert.Equal(t, c.(*context).echo, e)
	e.ReleaseContext(c)
}
