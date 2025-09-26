package fiber

import (
	"errors"
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestServerErrorHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := &App{}
	fctx := &fasthttp.RequestCtx{}
	err := errors.New("test error")

	// Act
	app.serverErrorHandler(fctx, err)

	// Assert
	assert.Equal(t, fasthttp.StatusInternalServerError, fctx.Response.StatusCode())
}
