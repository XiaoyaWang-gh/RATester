package fiber

import (
	"fmt"
	"testing"
)

func TestXHR_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		app: &App{
			getBytes: func(s string) (b []byte) {
				return []byte(s)
			},
		},
	}

	ctx.Set(HeaderXRequestedWith, "xmlhttprequest")

	if !ctx.XHR() {
		t.Error("Expected XHR to be true")
	}

	ctx.Set(HeaderXRequestedWith, "notxmlhttprequest")

	if ctx.XHR() {
		t.Error("Expected XHR to be false")
	}
}
