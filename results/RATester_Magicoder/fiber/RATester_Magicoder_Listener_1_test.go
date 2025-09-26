package fiber

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	cfg := ListenConfig{
		GracefulContext: context.Background(),
	}

	err = app.Listener(ln, cfg)
	if err != nil {
		t.Fatal(err)
	}
}
