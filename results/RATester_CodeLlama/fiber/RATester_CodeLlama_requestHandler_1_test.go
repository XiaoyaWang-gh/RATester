package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestRequestHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := &App{}
	rctx := &fasthttp.RequestCtx{}

	// Act
	app.requestHandler(rctx)

	// Assert
	// TODO: Add assertions
}
