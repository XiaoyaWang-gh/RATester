package gin

import (
	"fmt"
	"testing"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func TestHandler_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{}
	engine.UseH2C = true
	h2s := &http2.Server{}
	handler := h2c.NewHandler(engine, h2s)
	if handler == nil {
		t.Errorf("handler is nil")
	}
}
