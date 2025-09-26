package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestStale_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	c.fasthttp = &fasthttp.RequestCtx{}
	c.method = "GET"
	c.path = "/"
	c.pathOriginal = "/"
	c.treePath = "/"
	c.detectionPath = "/"
	c.indexRoute = 0
	c.indexHandler = 0
	c.methodINT = 1
	c.matched = true
	if c.Stale() {
		t.Errorf("Stale() should return false")
	}
}
