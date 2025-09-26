package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestserverErrorHandler_1(t *testing.T) {
	app := &App{}

	tests := []struct {
		name string
		app  *App
		fctx *fasthttp.RequestCtx
		err  error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			app.serverErrorHandler(tt.fctx, tt.err)
		})
	}
}
