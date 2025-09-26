package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestShutdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		server: &fasthttp.Server{},
	}

	err := app.Shutdown()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
