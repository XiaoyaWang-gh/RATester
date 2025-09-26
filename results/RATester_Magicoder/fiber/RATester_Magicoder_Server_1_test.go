package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		server: &fasthttp.Server{},
	}

	expected := app.server
	actual := app.Server()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
