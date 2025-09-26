package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestScheme_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
		app: &App{
			getString: func(b []byte) string {
				return string(b)
			},
		},
	}

	ctx.fasthttp.Request.Header.Set("X-Forwarded-Proto", "https")
	ctx.fasthttp.Request.Header.Set("X-Forwarded-Ssl", "on")
	ctx.fasthttp.Request.Header.Set("X-Url-Scheme", "http")

	scheme := ctx.Scheme()
	if scheme != "https" {
		t.Errorf("Expected scheme to be 'https', but got '%s'", scheme)
	}

	ctx.fasthttp.Request.Header.Set("X-Forwarded-Proto", "http")
	ctx.fasthttp.Request.Header.Set("X-Forwarded-Ssl", "off")
	ctx.fasthttp.Request.Header.Set("X-Url-Scheme", "https")

	scheme = ctx.Scheme()
	if scheme != "https" {
		t.Errorf("Expected scheme to be 'https', but got '%s'", scheme)
	}

	ctx.fasthttp.Request.Header.Set("X-Forwarded-Proto", "http")
	ctx.fasthttp.Request.Header.Set("X-Forwarded-Ssl", "off")
	ctx.fasthttp.Request.Header.Set("X-Url-Scheme", "http")

	scheme = ctx.Scheme()
	if scheme != "http" {
		t.Errorf("Expected scheme to be 'http', but got '%s'", scheme)
	}
}
