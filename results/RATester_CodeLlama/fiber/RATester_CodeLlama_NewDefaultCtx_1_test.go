package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNewDefaultCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	ctx := NewDefaultCtx(app)
	assert.Equal(t, app, ctx.app)
}
