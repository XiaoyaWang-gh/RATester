package fiber

import (
	"errors"
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestErrorHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		mountFields: &mountFields{
			appList: map[string]*App{
				"": {
					config: Config{
						ErrorHandler: func(ctx Ctx, err error) error {
							return ctx.Status(500).SendString("Internal Server Error")
						},
					},
				},
			},
		},
		config: Config{
			ErrorHandler: func(ctx Ctx, err error) error {
				return ctx.Status(500).SendString("Global Error Handler")
			},
		},
	}

	ctx := &DefaultCtx{
		app:      app,
		fasthttp: &fasthttp.RequestCtx{},
	}

	err := errors.New("test error")

	err = app.ErrorHandler(ctx, err)
	if err.Error() != "Internal Server Error" {
		t.Errorf("Expected error message to be 'Internal Server Error', but got '%s'", err.Error())
	}

	err = app.ErrorHandler(ctx, err)
	if err.Error() != "Global Error Handler" {
		t.Errorf("Expected error message to be 'Global Error Handler', but got '%s'", err.Error())
	}
}
