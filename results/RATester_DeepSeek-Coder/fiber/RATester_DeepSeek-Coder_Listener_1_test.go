package fiber

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		server: &fasthttp.Server{},
	}

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	cfg := ListenConfig{
		GracefulContext: context.Background(),
	}

	err = app.Listener(ln, cfg)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
