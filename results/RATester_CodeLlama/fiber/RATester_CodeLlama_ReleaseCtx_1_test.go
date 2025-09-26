package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestReleaseCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	c.release()
	app.pool.Put(c)
	assert.Equal(t, 1, app.pool.New)
	assert.Equal(t, 1, app.pool.Put)
}
