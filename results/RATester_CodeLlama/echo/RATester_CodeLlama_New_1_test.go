package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNew_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	assert.NotNil(t, e)
	assert.NotNil(t, e.filesystem)
	assert.NotNil(t, e.Server)
	assert.NotNil(t, e.TLSServer)
	assert.NotNil(t, e.AutoTLSManager)
	assert.NotNil(t, e.Logger)
	assert.NotNil(t, e.colorer)
	assert.NotNil(t, e.maxParam)
	assert.NotNil(t, e.ListenerNetwork)
}
