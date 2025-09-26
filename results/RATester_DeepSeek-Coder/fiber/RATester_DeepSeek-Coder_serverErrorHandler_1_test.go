package fiber

import (
	"errors"
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestServerErrorHandler_1(t *testing.T) {
	t.Parallel()

	app := &App{}

	tests := []struct {
		name    string
		fctx    *fasthttp.RequestCtx
		err     error
		wantErr bool
	}{
		{
			name:    "Test case 1",
			fctx:    &fasthttp.RequestCtx{},
			err:     errors.New("test error"),
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			app.serverErrorHandler(tt.fctx, tt.err)
			if (tt.fctx.Response.StatusCode() != StatusInternalServerError) != tt.wantErr {
				t.Errorf("serverErrorHandler() = %v, want %v", tt.fctx.Response.StatusCode(), tt.wantErr)
			}
		})
	}
}
