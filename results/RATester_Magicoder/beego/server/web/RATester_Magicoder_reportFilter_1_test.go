package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestreportFilter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{
			enableFilter: true,
			filters: [5][]*FilterRouter{
				{
					{
						filterFunc: func(ctx *context.Context) {
							// filter logic
						},
						pattern: "/test",
					},
				},
			},
		},
	}

	result := app.reportFilter()

	if len(result) == 0 {
		t.Error("Expected result to not be empty")
	}
}
